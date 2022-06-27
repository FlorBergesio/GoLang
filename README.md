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

## Sentencia for
- No hace falta usar los paréntesis a menos que sea necesario agrupar condiciones
- Reemplaza a todas las otras sentencias de ciclos disponibles en otro lenguajes
- No hace falta inicializar la variable i en el ciclo, se puede inicializar afuera
- No hace falta agregar un incremento a i en el ciclo, se puede agregar en el cuerpo del ciclo
- No hace falta agregar una condición de checkeo en el ciclo, se puede dejar vacío y salir del mismo con un break
- Por ejemplo, puede haber un for {...} que se ejecutará indefinidamente, hasta que algo dentro del cuerpo llege a una sentencia break
- Se puede utilizar la sentencia continue para mandar la ejecución al inicio del ciclo for en una nueva iteracion

## Sentencia goto
- Se puede definir un inicio de bloque utilizando un nombre en mayusculas y dos puntos. Ejemplo -> RUTINA:
- Se puede redireccionar el flujo de ejecución a este bloque utilizando la sentencia goto RUTINA

## Funciones
- Estructura: func nombre_funcion(nombre_parametro tipo_de_dato, nombre_parametro tipo_de_dato) (tipo_de_dato_return, tipo_de_dato_return) {... return valor, valor}
- No hace falta usar los parentesis en el tipo de dato de retorno si solo es 1
- Podemos recuperar los datos de return de esta forma: var1, var2 := nombre_funcion(parametros)
- Se puede definir un nombre de variable para un valor de retorno, en ese caso al escribir unicamente return se retornará el valor actual de dicha variable
- No existe la sobrecarga de funciones en Go
- Podemos definir que una funcion puede recibir un numero indefinido de parametros, usando la funcion range y el "_" para el valor de retorno que no nos interesa: suma(numero ...int) {
    for _, num := range numero{...}
}

## Funciones anónimas
- Se pueden definir variables del tipo func que contengan funciones anónimas, las cuales pueden cambiar en tiempo de ejecución
- Se debe respetar la definicion inicial de parametros y valores de return

### Closures
- Pueden acceder a variables creadas fuera de la función (en la función "original")
- Si se ejecuta varias veces la función, el código dentro de la funcion original pero fuera del closure solo se ejecutará la primera vez

## Vectores
- Los vectores no pueden variar sus dimensiones en tiempo real

### Slices
- Vectores que pueden variar sus dimensiones en tiempo real
- Se definen al no indicar la longitud en los corchetes. Ej: slice:= []int
- También se pueden crear utilizando la función make, la cual acepta los parámetros de longitud inicial y maximo de longitud reservada en memorita. Ej: slice := make([]int,5,20)
- El largo de un slice se puede obtener con len(nombreSlice), mientras que la capacidad máxima reservada en memoria se puede obtener con cap(nombreSlice)
- Se pueden agregar elementos mas alla de la capacidad definida, pero deberá hacerse con append(nombreSlice, elemento). En este caso si no se define un cap, se definirá automaticamente segun una potencia de 2

## Librerias utiles
- fmt: permite mostrar textos por pantalla
- os: permite manejar cuestiones del sistema operativo
- bufio: permite aceptar entradas por teclado

### Paquete fmt - verbos
- %d - numérico, base 10