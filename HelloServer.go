package main
import (
	"encoding/json"
	"net"
	"log"
	)
type ThreeTypeStructure struct {
	Head string //byte slice
	MidSeparatorTypes []byte
	Tail string
	EndTypes []byte

	}
	
func main () {
	
	var Hello ThreeTypeStructure 
	
	Hello.Head =  "Hello" 
	
	Hello.MidSeparatorTypes = []byte( " ,:" )
	
	Hello.Tail = "World"
	
	Hello.EndTypes = []byte ( "!.:" )
	
	
	byteSent , err := json.Marshal ( Hello )

	errChk ( err )

	tcpListen , err := net.Listen ( "tcp" , ":3779" )
	
	tcpAccept , err  := tcpListen.Accept ()

	errChk ( err ) 

	tcpAccept.Write ( byteSent )

	tcpAccept.Close ()
	
	tcpListen.Close ()
	
}

func errChk ( err error ) {

	if err != nil {
		
		log.Fatal ( err ) 

		}

}
