package btree

type BTree struct {
	size  int
	order int
	root  *bNode
	hot   *bNode
}

func (b *BTree) search(e int) *bNode {
	n := b.root
	b.hot = nil

	for n != nil {
		r := n.find(e)
		if r < len(n.key) && n.key[r] == e {
			return n
		}
		b.hot = n
		n = n.child[r]
	}
	return nil
}

func (b *BTree) Insert(e int) bool {
	n := b.search(e)
	if n != nil {
		return false
	}

	r := b.hot.find(e)
	tmp := b.hot.key[:r]
	tmp = append(tmp, e)
	b.hot.key = append(tmp, b.hot.key[r:]...)

	tmp1 := b.hot.child[0 : r+1]
	tmp1 = append(tmp1, nil)
	b.hot.child = append(tmp1, b.hot.child[r+1:]...)
	b.size++
	b.solveOverflow(b.hot)
	return true
}

func (b *BTree) Remove(e int) bool {
	n := b.search(e)
	if n == nil {
		return false
	}

	r := n.find(e)
	if n.child[0] != nil { // is not a leaf
		u := n.child[r+1]
		for u.child[0] != nil {
			u = u.child[0]
		}
		n.key[r] = u.key[0]
		n = u
		r = 0
	}
	tmp := n.key[:r]
	n.key = append(tmp, n.key[r+1:len(n.key)]...)

	tmp1 := n.child[:r+1]
	n.child = append(tmp1, n.child[r+2:len(n.child)]...)
	b.size--
	b.solveUnderflow(n)
	return true
}

func (b *BTree) solveUnderflow(n *bNode) {
}

func (b *BTree) solveOverflow(n *bNode) {
	if b.order >= len(n.child) {
		return
	}
	s := b.order / 2
	u := newNode()
	// move node in n to u
	u.key = append(u.key, n.key[s:]...)
	u.child = append(u.child, n.child[s:]...)
	// remove default nil node
	u.child = u.child[1:]

	// remove moved node
	n.key = n.key[:s]
	n.child = n.child[:s]
	if u.child[0] != nil {
		for j := 0; j < len(u.child); j++ {
			u.child[j].parent = u
		}
	}

	p := n.parent
	if p == nil {
		p = newNode()
		b.root = p
		p.child[0] = n
		n.parent = p
	}

	r := p.find(n.key[0])
	tmp := p.key[:r]
	tmp = append(tmp, n.key[len(n.key)-1])
	p.key = append(tmp, p.key[r:]...)

	temp1 := p.child[:r+1]
	temp1 = append(temp1, u)
	p.child = append(temp1, p.child[r+1:]...)

	u.parent = p
	b.solveOverflow(p)
}

func NewBTree(e ...int) *BTree {
	arg := []int{3}
	copy(arg, e)
	return &BTree{
		size:  0,
		order: arg[0],
		root:  newNode(),
		hot:   nil,
	}
}
