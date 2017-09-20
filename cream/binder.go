package cream

func RegisterBinder(method METHOD, mime MIME, binder Binder) {
	if _, ok := binders[method]; !ok {
		binders[method] = map[MIME]Binder{}
	}

	binders[method][mime] = binder
}

func binder(method METHOD, mime MIME) Binder {
	if bs, ok := binders[method]; ok {
		if b, ok := bs[mime]; ok {
			return b
		}
	}

	return defaultFormBinder
}

var binders = map[METHOD]map[MIME]Binder{
	METHODGet: map[MIME]Binder{
		MIMEApplicationForm: defaultFormBinder,
	},
}

type Binder interface {
	Bind() error
}

var defaultFormBinder Binder
