package main

import(
	"strconv"
	"fmt"
) 

func main() {

	var bloque datos_del_bloque
	bloque.hora_de_creacion_del_bloque=1020
	c:=bloque.string_para_hora_de_creacion_del_bloque()
	cc:=bloque.pasar_a_bytes_hora_de_creacion_del_bloque(c)
	fmt.Println(cc)
}

type datos_del_bloque struct{
	hora_de_creacion_del_bloque int64
	hash_del_bloque_anterior []byte
	hash_del_bloque_actual []byte
	datos_del_bloque []byte  
}

func (b *datos_del_bloque) string_para_hora_de_creacion_del_bloque () string{
	bb:= strconv.FormatInt(b.hora_de_creacion_del_bloque, 10)
	return bb
}

func (bloque *datos_del_bloque) pasar_a_bytes_hora_de_creacion_del_bloque (bb string) []byte{
	fmt.Println(bb)
	fmt.Println("Pasa por aqui")
	r:= []byte(bb)
	return r
}