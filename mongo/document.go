package mongo

import (
	"github.com/ungerik/go-start/mgo/bson"
	"github.com/ungerik/go-start/model"
)

///////////////////////////////////////////////////////////////////////////////
// Document

type Document interface {
	Init(collection *Collection, embeddingStruct interface{})
	Collection() *Collection
	ObjectId() bson.ObjectId
	SetObjectId(id bson.ObjectId)
	Iterator() model.Iterator
	Ref() Ref
	Save() error
	Remove() error
}
