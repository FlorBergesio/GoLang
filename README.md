# GoLang

## Notas
- Se puede correr un desarrollo en go con go run main.go
- Se puede buildear un desarrollo en go con go build main.go
- El ejecutable de go tiene todo lo necesario para poder ejecutarse desde cualquier sistema operativo
- Las funciones en go deben abrir llaves en la misma linea donde se define el nombre

## Variables
Existen 3 tipos de variables:
- Numerics: Enteras (int, int8, etc), enteras sin signo (uint, uint8, etc), punto flotante (float32, float64, etc). Se inicializan con 0
- Strings: (string). Se inicializan con cadena vacía
- Booleans: (bool). Se inicializan con false
Para definir una variable se utilizan las sintaxis: 
- var var1, var2, var3 int
- var nombre_variable tipo_variable = valor_variable
- nombre_variable := valor variable
- var1, var2, var3 := 5
- var1, var2, var3 := 1, 2, "Texto"
Parseo:
- Los numeros se pueden parsear a numeros enteros de la siguiente forma: int(variable_valor_int64)
- Los string necesitan el uso de paquetes: fmt.Sprintf("%d", variable_valor_int64) o strconv.Itoa(int(variable_valor_int64))

## Scope
El scope de las variables, métodos y funciones se determina de la siguiente manera:
- Si el nombre comienza con una letra minuscula, es privado
- Si el nombre comienza con una letra mayuscula, se exporta a otros scope (paquetes, desarrollos, etc)
- Si una variable se define fuera de las funciones, puede ser utilizada por todas las funciones del paquete
- Si una variable se define dentro de esa funcion, su scope se limita a esa funcion

## Paquete main
- El archivo principal de un desarrollo en go es main.go
- Debe tener la definicion de package main
- La funcion main es la entrada al programa

## Sentencia if
- No hace falta usar los paréntesis a menos que sea necesario agrupar condiciones
- Se pueden asignar variables dentro de la sentencia if y estas variables serán accesibles en el cuerpo del if y else: if estado=false; estado == true {...}

## Sentencia switch
- No hace falta usar los paréntesis a menos que sea necesario agrupar condiciones
- Se pueden asignar variables dentro de la sentencia if y estas variables serán accesibles en el cuerpo del switch: switch numero := 5; numero { case 1: ... }
- No hace falta utilizar el break

## Librerias utiles
- fmt: permite mostrar textos por pantalla
- os: permite manejar cuestiones del sistema operativo
- bufio: permite aceptar entradas por teclado

### Paquete fmt - verbos
- %d - numérico, base 10