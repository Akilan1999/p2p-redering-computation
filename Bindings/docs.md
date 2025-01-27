<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# Bindings

```go
import "github.com/Akilan1999/p2p-rendering-computation/Bindings"
```

## Index

- [func ConvertStructToJSONString\(Struct interface\{\}\) \*C.char](<#ConvertStructToJSONString>)
- [func EscapeFirewall\(HostOutsideNATIP string, HostOutsideNATPort string, internalPort string\) \(output \*C.char\)](<#EscapeFirewall>)
- [func GetSpecs\(IP string\) \(output \*C.char\)](<#GetSpecs>)
- [func Init\(customConfig string\) \(output \*C.char\)](<#Init>)
- [func MapPort\(Port string, DomainName string, ServerAddress string\) \*C.char](<#MapPort>)
- [func RemoveContainer\(IP string, ID string\) \(output \*C.char\)](<#RemoveContainer>)
- [func Server\(\) \(output \*C.char\)](<#Server>)
- [func StartContainer\(IP string\) \(output \*C.char\)](<#StartContainer>)
- [func UpdateIPTable\(\) \(output \*C.char\)](<#UpdateIPTable>)
- [func ViewIPTable\(\) \(output \*C.char\)](<#ViewIPTable>)


<a name="ConvertStructToJSONString"></a>
## func [ConvertStructToJSONString](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L159>)

```go
func ConvertStructToJSONString(Struct interface{}) *C.char
```



<a name="EscapeFirewall"></a>
## func [EscapeFirewall](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L118>)

```go
func EscapeFirewall(HostOutsideNATIP string, HostOutsideNATPort string, internalPort string) (output *C.char)
```



<a name="GetSpecs"></a>
## func [GetSpecs](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L79>)

```go
func GetSpecs(IP string) (output *C.char)
```



<a name="Init"></a>
## func [Init](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L88>)

```go
func Init(customConfig string) (output *C.char)
```



<a name="MapPort"></a>
## func [MapPort](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L137>)

```go
func MapPort(Port string, DomainName string, ServerAddress string) *C.char
```



<a name="RemoveContainer"></a>
## func [RemoveContainer](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L29>)

```go
func RemoveContainer(IP string, ID string) (output *C.char)
```



<a name="Server"></a>
## func [Server](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L148>)

```go
func Server() (output *C.char)
```



<a name="StartContainer"></a>
## func [StartContainer](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L20>)

```go
func StartContainer(IP string) (output *C.char)
```



<a name="UpdateIPTable"></a>
## func [UpdateIPTable](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L109>)

```go
func UpdateIPTable() (output *C.char)
```



<a name="ViewIPTable"></a>
## func [ViewIPTable](<https://github.com/Akilan1999/p2p-rendering-computation/blob/master/Bindings/Client.go#L100>)

```go
func ViewIPTable() (output *C.char)
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)