package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"

	"github.com/moutend/go-wca"
)

func main() {
	// IPropertyStore 886d8eeb-8cf2-4446-8d02-cdba1dbdcf99
	var guid = syscall.GUID{
		Data1: 0x886d8eeb,
		Data2: 0x8cf2,
		Data3: 0x4446,
		Data4: [8]byte{0x8d, 0x02, 0xcd, 0xba, 0x1d, 0xbd, 0xcf, 0x99},
	}

	ole.CoInitialize(0)
	var ps *wca.IPropertyStore
	var shell32 = syscall.NewLazyDLL("shell32.dll")
	var procWriteConsoleW = shell32.NewProc("SHGetPropertyStoreFromParsingName")
	ret, _, err := procWriteConsoleW.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:\\dev\\go_work\\go-property\\w.mp4"))),
		0, 2,
		uintptr(unsafe.Pointer(&guid)), uintptr(unsafe.Pointer(&ps)))
	fmt.Println(ret, err)
	fmt.Printf("%v\n", ps)

	// var k uint32
	// var pk *wca.PropertyKey
	// p.GetAt(k, pk)
	// fmt.Println(pk.String())
	var count uint32
	ps.GetCount(&count)
	fmt.Println(count)

	var propsys = syscall.NewLazyDLL("propsys.dll")
	var procPSGetNameFromPropertyKey = propsys.NewProc("PSGetNameFromPropertyKey")
	var procPSFormatForDisplayAlloc = propsys.NewProc("PSFormatForDisplayAlloc")
	PDFF_DEFAULT := 0x00000000

	var procPSGetPropertyKeyFromName = propsys.NewProc("PSGetPropertyKeyFromName")

	var i uint32
	for i = 0; i < count; i++ {
		var key wca.PropertyKey
		ps.GetAt(i, &key)

		// fmt.Println(key.GUID)
		// fmt.Println(key.PID)

		var pv wca.PROPVARIANT
		ps.GetValue(&key, &pv)
		fmt.Println(pv.VT)

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
		ret2, _, _ := procPSFormatForDisplayAlloc.Call(
			uintptr(unsafe.Pointer(&key)), uintptr(unsafe.Pointer(&pv)),
			uintptr(unsafe.Pointer(&PDFF_DEFAULT)),
			uintptr(unsafe.Pointer(&propValue)))
		// fmt.Println(ret2, err)
		if ret2 == 0 {
			fmt.Printf("%s\n", syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(propValue))[:]))
		}

	}

	// PSGetPropertyKeyFromName
	name := "System.Media.SubTitle"
	var key wca.PropertyKey
	ret3, _, err := procPSGetPropertyKeyFromName.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(&key)))
	fmt.Println(ret3, err)

	var pv wca.PROPVARIANT
	ps.GetValue(&key, &pv)
	fmt.Println(pv.VT)

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
	ret2, _, _ := procPSFormatForDisplayAlloc.Call(
		uintptr(unsafe.Pointer(&key)), uintptr(unsafe.Pointer(&pv)),
		uintptr(unsafe.Pointer(&PDFF_DEFAULT)),
		uintptr(unsafe.Pointer(&propValue)))
	// fmt.Println(ret2, err)
	if ret2 == 0 {
		fmt.Printf("%s\n", syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(propValue))[:]))
	}

}
