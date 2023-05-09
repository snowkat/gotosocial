// Code generated by astool. DO NOT EDIT.

package streams

import (
	typepropertyvalue "github.com/superseriousbusiness/activity/streams/impl/schema/type_propertyvalue"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
)

// IsOrExtendsSchemaPropertyValue returns true if the other provided type is the
// PropertyValue type or extends from the PropertyValue type.
func IsOrExtendsSchemaPropertyValue(other vocab.Type) bool {
	return typepropertyvalue.IsOrExtendsPropertyValue(other)
}
