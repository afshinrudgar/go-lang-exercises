package erratum

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	for {
		if resource, err = o(); err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return
		}
	}
	defer resource.Close()
	defer func(resource Resource) {
		if r := recover(); r != nil {
			err = r.(error)
			if re, ok := r.(FrobError); ok {
				resource.Defrob(re.defrobTag)
			}
		}
	}(resource)
	resource.Frob(input)
	err = nil
	return
}
