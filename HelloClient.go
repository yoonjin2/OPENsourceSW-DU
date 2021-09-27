package main
import (
	"encoding/json"
	"net"
	"io"
	"bytes"
	"strconv"
	"log"
	"fmt"
	"bufio"
	"os"
	)
type FourTypeStructure struct {
	Head string //byte slice
	MidSeparatorTypes []byte
	Tail string
	EndTypes []byte

	}
	
func main () {
	
	if len ( os.Args ) < 4 {
	
		println ( "Please provide your server's ip address" )
	
		return
	
	}
	var Hello FourTypeStructure 
	
	tcpDial , err := net.Dial ( "tcp" , os.Args [ 1 ] + ":3779" ) 
	
	if err != nil {

		log.Fatal ( err ) 

	}

	
	reader := bufio.NewReader ( tcpDial )

	byteSent , err := io.ReadAll ( reader )

	byteSent = bytes.Trim ( byteSent , "\x00" )


	
	if err != nil {
	
	log.Fatal ( err ) 
	
	}

	err = json.Unmarshal ( byteSent , &Hello )
	
	if err != nil {
	
		log.Fatal ( err )
	 
	} 

	tcpDial.Close ()
	byteHello := make ( []byte , 13 ) 

	byteHello = append ( byteHello , []byte ( Hello.Head ) ... )

	index , err := strconv.Atoi ( os.Args [ 2 ] )  

	indexChk ( index , err )

	index --
	
	byteHello = append ( byteHello , Hello.MidSeparatorTypes [ index ] )

	byteHello = append ( byteHello , []byte ( Hello.Tail ) ... )
	
	index , err = strconv.Atoi ( os.Args [ 3 ] ) 

	indexChk ( index , err )

	index --

	byteHello = append ( byteHello , Hello.EndTypes [ index ] )

	fmt.Println ( string ( byteHello ) )


}

func indexChk ( index int , err error ) {

	if ( err != nil ) || (  index > 3 ) {

		os.Exit ( 1 ) 

	}

}

