package main

import (
	"fmt"
)
// 来自gin 框架
type HandlersChain []string

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
	catchAll
)

type node struct {
	// 节点路径，比如上面的s，earch，和upport
	path string
	// 节点是否是参数节点，比如上面的:post
	wildChild bool
	// 节点类型，包括static, root, param, catchAll
	// static: 静态节点，比如上面的s，earch等节点
	// root: 树的根节点
	// catchAll: 有*匹配的节点
	// param: 参数节点
	nType nodeType
	// 路径上最大参数个数
	maxParams uint8
	// 和children字段对应, 保存的是分裂的分支的第一个字符
	// 例如search和support, 那么s节点的indices对应的"eu"
	// 代表有两个分支, 分支的首字母分别是e和u
	indices string
	// 儿子节点
	children []*node
	// 处理函数
	handlers HandlersChain
	// 优先级，子节点注册的handler数量
	priority uint32
}

func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] == ':' || path[i] == '*' {
			n++
		}
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

// increments priority of the given child and reorders if necessary.
func (n *node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	// adjust position (move to front)
	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		// swap node positions
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	// build new index char string
	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++
	numParams := countParams(path)
	// non-empty tree
	if len(n.path) > 0 || len(n.children) > 0 {
	walk:
		for {
			// Update maxParams of the current node
			if numParams > n.maxParams {
				n.maxParams = numParams
			}
			// Find the longest common prefix.
			// This also implies that the common prefix contains no ':' or '*'
			// since the existing key can't contain those chars.
			i := 0
			max := min(len(path), len(n.path))
			for i < max && path[i] == n.path[i] {
				i++
			}
			// Split edge
			// 开始分裂，比如一开始path是search，新来了support，s是他们匹配的部分，
			// 那么会将s拿出来作为parent节点，增加earch和upport作为child节点
			if i < len(n.path) {
				child := node{
					path:      n.path[i:], // 不匹配的部分作为child节点
					wildChild: n.wildChild,
					indices:   n.indices,
					children:  n.children,
					handlers:  n.handlers,
					priority:  n.priority - 1, // 降级成子节点，priority减1
				}
				// Update maxParams (max of all children)
				for i := range child.children {
					if child.children[i].maxParams > child.maxParams {
						child.maxParams = child.children[i].maxParams
					}
				}

				// 当前节点的子节点变成刚刚分裂的出来的节点
				n.children = []*node{&child}
				// []byte for proper unicode char conversion, see #65
				n.indices = string([]byte{n.path[i]})
				n.path = path[:i]
				n.handlers = nil
				n.wildChild = false
			}
			// Make new node a child of this node
			// 将新来的节点插入新的parent节点作为子节点
			if i < len(path) {
				path = path[i:]
				// 如果是参数节点（包含:或*）
				if n.wildChild {
					n = n.children[0]
					n.priority++
					// Update maxParams of the child node
					if numParams > n.maxParams {
						n.maxParams = numParams
					}
					numParams--
					// Check if the wildcard matches
					// 例如：/blog/:pp 和 /blog/:ppp，需要检查更长的通配符
					if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
						// check for longer wildcard, e.g. :name and :names
						if len(n.path) >= len(path) || path[len(n.path)] == '/' {
							continue walk
						}
					}
					panic("path segment '" + path +
						"' conflicts with existing wildcard '" + n.path +
						"' in path '" + fullPath + "'")
				}
				// 首字母，用来与indices做比较
				c := path[0]
				// slash after param
				if n.nType == param && c == '/' && len(n.children) == 1 {
					n = n.children[0]
					n.priority++
					continue walk
				}
				// Check if a child with the next path byte exists
				// 判断子节点中是否有和当前path有匹配的，只需要查看子节点path的第一个字母即可，即indices
				// 比如s的子节点现在是earch和upport，indices为eu
				// 如果新来的路由为super，那么就是和upport有匹配的部分u，将继续分类现在的upport节点
				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}
				// Otherwise insert it
				if c != ':' && c != '*' {
					// []byte for proper unicode char conversion, see #65
					// 记录第一个字符，放在indices中
					n.indices += string([]byte{c})
					child := &node{
						maxParams: numParams,
					}
					// 增加子节点
					n.children = append(n.children, child)
					n.incrementChildPrio(len(n.indices) - 1)
					n = child
				}
				n.insertChild(numParams, path, fullPath, handlers)
				return
			} else if i == len(path) { // Make node a (in-path) leaf
				// 路径相同，如果已有handler就报错，没有就赋值
				if n.handlers != nil {
					panic("handlers are already registered for path ''" + fullPath + "'")
				}
				n.handlers = handlers
			}
			return
		}
	} else { // Empty tree，空树，插入节点，节点种类是root
		n.insertChild(numParams, path, fullPath, handlers)
		n.nType = root
	}
}

// @1: 参数个数
// @2: 路径
// @3: 完整路径
// @4: 处理函数
func (n *node) insertChild(numParams uint8, path string, fullPath string, handlers HandlersChain) {
	var offset int // already handled bytes of the path
	// find prefix until first wildcard (beginning with ':'' or '*'')
	// 找到前缀，只要匹配到wildcard
	for i, max := 0, len(path); numParams > 0; i++ {
		c := path[i]
		if c != ':' && c != '*' {
			continue
		}
		// find wildcard end (either '/' or path end)
		end := i + 1
		for end < max && path[end] != '/' {
			switch path[end] {
			// the wildcard name must not contain ':' and '*'
			case ':', '*':
				panic("only one wildcard per path segment is allowed, has: '" +
					path[i:] + "' in path '" + fullPath + "'")
			default:
				end++
			}
		}
		// check if this Node existing children which would be
		// unreachable if we insert the wildcard here
		if len(n.children) > 0 {
			panic("wildcard route '" + path[i:end] +
				"' conflicts with existing children in path '" + fullPath + "'")
		}
		// check if the wildcard has a name
		if end-i < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}
		if c == ':' { // param
			// split path at the beginning of the wildcard
			if i > 0 {
				n.path = path[offset:i]
				offset = i
			}
			child := &node{
				nType:     param,
				maxParams: numParams,
			}
			n.children = []*node{child}
			n.wildChild = true
			n = child
			n.priority++
			numParams--
			// if the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			if end < max {
				n.path = path[offset:end]
				offset = end

				child := &node{
					maxParams: numParams,
					priority:  1,
				}
				n.children = []*node{child}
				// 下次循环这个新的child节点
				n = child
			}
		} else { // catchAll
			if end != max || numParams > 1 {
				panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
			}
			if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
				panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
			}
			// currently fixed width 1 for '/'
			i--
			if path[i] != '/' {
				panic("no / before catch-all in path '" + fullPath + "'")
			}
			n.path = path[offset:i]
			// first node: catchAll node with empty path
			child := &node{
				wildChild: true,
				nType:     catchAll,
				maxParams: 1,
			}
			n.children = []*node{child}
			n.indices = string(path[i])
			n = child
			n.priority++
			// second node: node holding the variable
			child = &node{
				path:      path[i:],
				nType:     catchAll,
				maxParams: 1,
				handlers:  handlers,
				priority:  1,
			}
			n.children = []*node{child}
			return
		}
	}
	// insert remaining path part and handle to the leaf
	n.path = path[offset:]
	n.handlers = handlers
}

func main() {

	n := new(node)

	h1 := HandlersChain{"a"}
	n.addRoute("/v1", h1)
	h2 := HandlersChain{"b"}
	n.addRoute("/v1/foo", h2)
	h3 := HandlersChain{"c"}
	n.addRoute("/v1/foo/bar", h3)
	n.addRoute("/v1/foo/barxyz/:name", h3)

	for n != nil {
		fmt.Printf("%+v\n", n)
		if n.children != nil {
			n = n.children[0]
		} else {
			n = nil
		}
	}
}
