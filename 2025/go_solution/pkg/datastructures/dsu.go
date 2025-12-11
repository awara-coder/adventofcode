package datastructures

type DSU struct {
	capacity int
	parent   []int
	size     []int
}

func NewDSU(capacity int) *DSU {
	parent := make([]int, capacity)
	size := make([]int, capacity)
	for i := range capacity {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{
		capacity: capacity,
		parent:   parent,
		size:     size,
	}
}

func (d *DSU) Find(n int) int {
	if n >= d.capacity || n < 0 {
		return -1
	}

	if n == d.parent[n] {
		return n
	}

	d.parent[n] = d.Find(d.parent[n])
	return d.parent[n]
}

func (d *DSU) GetSize(n int) int {
	if n >= d.capacity || n < 0 {
		return -1
	}

	parentN := d.Find(n)

	return d.size[parentN]
}

func (d *DSU) Add(u, v int) bool {
	// If both are same, or parent is same, return false
	if u == v {
		return false
	}

	parentU := d.Find(u)
	parentV := d.Find(v)

	if parentU == parentV {
		return false
	}

	if d.size[parentU] < d.size[parentV] {
		parentU, parentV = parentV, parentU
	}

	// Merge them and return true
	d.parent[parentV] = parentU
	d.size[parentU] += d.size[parentV]

	return true
}
