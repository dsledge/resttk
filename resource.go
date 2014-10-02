package resttk

type ResourceInterface interface {
	FindAllResources(v interface{}) interface{}
	FindResource(key string, value string, v interface{}) interface{}
	SaveResource(key string, value string, v interface{}) interface{}
	UpdateResource(key string, value string, v interface{}) interface{}
	DeleteResource(key string, value string) bool
}

type BaseResource struct {
	parent ResourceInterface
}

func (r *BaseResource) Init(p ResourceInterface) {
	r.parent = p
}

func (r *BaseResource) FindAllResources(v interface{}) interface{} {
	resources := r.parent.FindAllResources(v)
	if resources != nil {
		return resources
	}

	return nil
}

func (r *BaseResource) FindResource(key string, value string, v interface{}) interface{} {
	resource := r.parent.FindResource(key, value, v)
	if resource != nil {
		return resource
	}

	return nil
}

func (r *BaseResource) SaveResource(key string, value string, v interface{}) interface{} {
	resource := r.parent.FindResource(key, value, v)
	if resource != nil {
		return resource
	}

	return nil
}

func (r *BaseResource) UpdateResource(key string, value string, v interface{}) interface{} {
	resource := r.parent.UpdateResource(key, value, v)
	if resource != nil {
		return resource
	}

	return nil
}

func (r *BaseResource) DeleteResource(key string, value string) bool {
	return r.parent.DeleteResource(key, value)
}
