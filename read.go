package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca"
)

// IPropertyStore 886d8eeb-8cf2-4446-8d02-cdba1dbdcf99
var guid = syscall.GUID{
	Data1: 0x886d8eeb,
	Data2: 0x8cf2,
	Data3: 0x4446,
	Data4: [8]byte{0x8d, 0x02, 0xcd, 0xba, 0x1d, 0xbd, 0xcf, 0x99},
}
var shell32 = syscall.NewLazyDLL("shell32.dll")
var procSHGetPropertyStoreFromParsingName = shell32.NewProc("SHGetPropertyStoreFromParsingName")
var propsys = syscall.NewLazyDLL("propsys.dll")
var procPSGetNameFromPropertyKey = propsys.NewProc("PSGetNameFromPropertyKey")
var procPSFormatForDisplayAlloc = propsys.NewProc("PSFormatForDisplayAlloc")
var procPSGetPropertyKeyFromName = propsys.NewProc("PSGetPropertyKeyFromName")
var pdffDefault int = 0x00000000

func main() {

	ole.CoInitialize(0)
	var ps *wca.IPropertyStore
	ret, _, err := procSHGetPropertyStoreFromParsingName.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:\\dev\\go_work\\go-property\\w.mp4"))),
		0, 2,
		uintptr(unsafe.Pointer(&guid)), uintptr(unsafe.Pointer(&ps)))
	fmt.Println(ret, err)
	fmt.Printf("%v\n", ps)

	var count uint32
	ps.GetCount(&count)
	fmt.Println(count)

	// get all property key and value from file
	var i uint32
	for i = 0; i < count; i++ {
		var key wca.PropertyKey
		ps.GetAt(i, &key)

		// fmt.Println(key.GUID)
		// fmt.Println(key.PID)

		var pv wca.PROPVARIANT
		ps.GetValue(&key, &pv)
		// fmt.Println(pv.VT)

		// PSGetNameFromPropertyKey
		var propKey uintptr
		ret, _, _ := procPSGetNameFromPropertyKey.Call(
			uintptr(unsafe.Pointer(&key)),
			uintptr(unsafe.Pointer(&propKey)))
		// fmt.Println(ret, err)
		if ret == 0 {
			fmt.Printf("%s=", syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(propKey))[:]))
		}

		// PSFormatForDisplayAlloc
		var propValue uintptr
		ret, _, _ = procPSFormatForDisplayAlloc.Call(
			uintptr(unsafe.Pointer(&key)), uintptr(unsafe.Pointer(&pv)),
			uintptr(unsafe.Pointer(&pdffDefault)),
			uintptr(unsafe.Pointer(&propValue)))
		// fmt.Println(ret2, err)
		if ret == 0 {
			fmt.Printf("%s\n", syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(propValue))[:]))
		}

	}

	// get single property key and value from file
	// PSGetPropertyKeyFromName
	name := "System.Media.SubTitle"
	var key wca.PropertyKey
	ret, _, err = procPSGetPropertyKeyFromName.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(&key)))
	// fmt.Println(ret, err)

	var pv wca.PROPVARIANT
	ps.GetValue(&key, &pv)
	// fmt.Println(pv.VT)

	// PSGetNameFromPropertyKey
	var propKey uintptr
	ret, _, _ = procPSGetNameFromPropertyKey.Call(
		uintptr(unsafe.Pointer(&key)),
		uintptr(unsafe.Pointer(&propKey)))
	// fmt.Println(ret, err)
	if ret == 0 {
		fmt.Printf("%s=", syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(propKey))[:]))
	}

	// PSFormatForDisplayAlloc
	var propValue uintptr
	ret, _, _ = procPSFormatForDisplayAlloc.Call(
		uintptr(unsafe.Pointer(&key)), uintptr(unsafe.Pointer(&pv)),
		uintptr(unsafe.Pointer(&pdffDefault)),
		uintptr(unsafe.Pointer(&propValue)))
	// fmt.Println(ret, err)
	if ret == 0 {
		fmt.Printf("%s\n", syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(propValue))[:]))
	}
}
