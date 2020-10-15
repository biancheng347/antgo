package container

type Array struct {
	Slice []interface{}
}

// New ...
func NewArray() *Array {
	return &Array{Slice: make([]interface{}, 0)}
}

// Set ...
func (this *Array) Append(value interface{}) {
	this.Slice = append(this.Slice, value)
}

// Count ...
func (this *Array) Len() int {
	return len(this.Slice)
}

// List ...
func (this *Array) List() []interface{} {
	return this.Slice
}

// InsertAfter ...
func (this *Array) InsertAfter(index int, value interface{}) []interface{} {
	var reset = make([]interface{}, 0)
	prefix := append(reset, this.Slice[index:]...)
	this.Slice = append(this.Slice[0:index], value)
	this.Slice = append(this.Slice, prefix...)
	return this.Slice
}

// Delete ...
func (this *Array) Delete(index int) []interface{} {
	this.Slice = append(this.Slice[:index], this.Slice[index+1:]...)
	return this.Slice
}

// Set ...
func (this *Array) Set(index int, value interface{}) {
	this.Slice[index] = value
}

// Set ...
func (this *Array) Get(index int) interface{} {
	return this.Slice[index]
}

// Search ...
func (this *Array) Search(value interface{}) int {
	for i := 0; i < len(this.Slice); i++ {
		if this.Slice[i] == value {
			return i
		}
	}
	return 0
}

// Search ...
func (this *Array) Clear() {
	this.Slice = make([]interface{}, 0)
}
