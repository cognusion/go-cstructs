

# cstructs
`import "github.com/cognusion/go-cstructs"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type CSlice](#CSlice)
  * [func (c *CSlice) Append(thing interface{})](#CSlice.Append)
  * [func (c *CSlice) Cap() int](#CSlice.Cap)
  * [func (c *CSlice) Get(index int) interface{}](#CSlice.Get)
  * [func (c *CSlice) GetSlice() []interface{}](#CSlice.GetSlice)
  * [func (c *CSlice) InitSlice(size int)](#CSlice.InitSlice)
  * [func (c *CSlice) Insert(index int, thing interface{})](#CSlice.Insert)
  * [func (c *CSlice) Len() int](#CSlice.Len)
* [type CStringMapSlice](#CStringMapSlice)
  * [func (c *CStringMapSlice) Append(thing map[string]interface{})](#CStringMapSlice.Append)
  * [func (c *CStringMapSlice) Cap() int](#CStringMapSlice.Cap)
  * [func (c *CStringMapSlice) Get(index int) map[string]interface{}](#CStringMapSlice.Get)
  * [func (c *CStringMapSlice) GetSlice() []map[string]interface{}](#CStringMapSlice.GetSlice)
  * [func (c *CStringMapSlice) InitSlice(size int)](#CStringMapSlice.InitSlice)
  * [func (c *CStringMapSlice) Insert(index int, thing map[string]interface{})](#CStringMapSlice.Insert)
  * [func (c *CStringMapSlice) Len() int](#CStringMapSlice.Len)


#### <a name="pkg-files">Package files</a>
[cslice.go](/src/github.com/cognusion/go-cstructs/cslice.go) [cstringmapslice.go](/src/github.com/cognusion/go-cstructs/cstringmapslice.go) 






## <a name="CSlice">type</a> [CSlice](/src/target/cslice.go?s=81:139#L8)
``` go
type CSlice struct {
    sync.RWMutex
    // contains filtered or unexported fields
}

```
CSlice is a goro-safe interface{} slice










### <a name="CSlice.Append">func</a> (\*CSlice) [Append](/src/target/cslice.go?s=1267:1309#L59)
``` go
func (c *CSlice) Append(thing interface{})
```
Append s the specified item onto the underlying slice




### <a name="CSlice.Cap">func</a> (\*CSlice) [Cap](/src/target/cslice.go?s=193:219#L14)
``` go
func (c *CSlice) Cap() int
```
Cap returns the capacity of the underlying slice




### <a name="CSlice.Get">func</a> (\*CSlice) [Get](/src/target/cslice.go?s=962:1005#L45)
``` go
func (c *CSlice) Get(index int) interface{}
```
Get returns the value of the interface{} at the specific index.
Care should be taken that 0 =< index < size, else panic




### <a name="CSlice.GetSlice">func</a> (\*CSlice) [GetSlice](/src/target/cslice.go?s=1118:1159#L52)
``` go
func (c *CSlice) GetSlice() []interface{}
```
GetSlice returns the underlying interface{} slice




### <a name="CSlice.InitSlice">func</a> (\*CSlice) [InitSlice](/src/target/cslice.go?s=517:553#L29)
``` go
func (c *CSlice) InitSlice(size int)
```
InitSlice [re]initializes the underlying slice with the specified size,
enabling safe Insert operations




### <a name="CSlice.Insert">func</a> (\*CSlice) [Insert](/src/target/cslice.go?s=730:783#L37)
``` go
func (c *CSlice) Insert(index int, thing interface{})
```
Insert should be used only after InitSlice is used, and care is taken
that 0 =< index < size, else panic




### <a name="CSlice.Len">func</a> (\*CSlice) [Len](/src/target/cslice.go?s=325:351#L21)
``` go
func (c *CSlice) Len() int
```
Len returns the length of the underlying slice




## <a name="CStringMapSlice">type</a> [CStringMapSlice](/src/target/cstringmapslice.go?s=101:179#L8)
``` go
type CStringMapSlice struct {
    sync.RWMutex
    // contains filtered or unexported fields
}

```
CStringMapSlice is a goro-safe map[string]interface{} slice










### <a name="CStringMapSlice.Append">func</a> (\*CStringMapSlice) [Append](/src/target/cstringmapslice.go?s=1405:1467#L59)
``` go
func (c *CStringMapSlice) Append(thing map[string]interface{})
```
Append s the specified item onto the underlying slice




### <a name="CStringMapSlice.Cap">func</a> (\*CStringMapSlice) [Cap](/src/target/cstringmapslice.go?s=233:268#L14)
``` go
func (c *CStringMapSlice) Cap() int
```
Cap returns the capacity of the underlying slice




### <a name="CStringMapSlice.Get">func</a> (\*CStringMapSlice) [Get](/src/target/cstringmapslice.go?s=1060:1123#L45)
``` go
func (c *CStringMapSlice) Get(index int) map[string]interface{}
```
Get returns the value of the interface{} at the specific index.
Care should be taken that 0 =< index < size, else panic




### <a name="CStringMapSlice.GetSlice">func</a> (\*CStringMapSlice) [GetSlice](/src/target/cstringmapslice.go?s=1236:1297#L52)
``` go
func (c *CStringMapSlice) GetSlice() []map[string]interface{}
```
GetSlice returns the underlying interface{} slice




### <a name="CStringMapSlice.InitSlice">func</a> (\*CStringMapSlice) [InitSlice](/src/target/cstringmapslice.go?s=575:620#L29)
``` go
func (c *CStringMapSlice) InitSlice(size int)
```
InitSlice [re]initializes the underlying slice with the specified size,
enabling safe Insert operations




### <a name="CStringMapSlice.Insert">func</a> (\*CStringMapSlice) [Insert](/src/target/cstringmapslice.go?s=808:881#L37)
``` go
func (c *CStringMapSlice) Insert(index int, thing map[string]interface{})
```
Insert should be used only after InitSlice is used, and care is taken
that 0 =< index < size, else panic




### <a name="CStringMapSlice.Len">func</a> (\*CStringMapSlice) [Len](/src/target/cstringmapslice.go?s=374:409#L21)
``` go
func (c *CStringMapSlice) Len() int
```
Len returns the length of the underlying slice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
