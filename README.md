# go-property
IPropertyStore 를 이용한 파일 프로퍼티 취득, 등록  
  
파일 프로퍼티를 읽어들일 수 있는 prop.exe 가 한글, 일본어에 대응하지 않아 문자가 깨져서 표시 되기 때문에 golang 에서 다시 작성한다.  
Property System Command Line Interface (Prop.exe) - https://archive.codeplex.com/?p=prop
  
사용 DLL, function  
--
shell32.dll
* SHGetPropertyStoreFromParsingName  

propsys.dll  
* PSGetNameFromPropertyKey  
* PSFormatForDisplayAlloc  
* PSGetPropertyKeyFromName  
  
목표
--
prop.exe 의 기능 중에서 dump, set 의 기능을 구현한다.  
dump : 파일이 가진 모든 프로퍼티의 키와 값을 출력한다.  
set : 프로퍼티명으로 값을 입력한다.  (set은 답이 없는듯... ㅠㅠ)
  
  
