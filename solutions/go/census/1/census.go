// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name: name,
		Age: age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return len(r.Name) > 0 && r.Address != nil &&  len(r.Address["street"]) > 0
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	var cleanResident Resident
	r.Name = cleanResident.Name
	r.Age = cleanResident.Age
	r.Address = cleanResident.Address
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var count int
	for _,v := range residents {
		if v.HasRequiredInfo() {
			count++
		}
	}
	return count
}
