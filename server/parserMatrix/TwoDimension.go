package parsermatrix

func ConvertirMatrizToString(matrizInterface [][]interface{}) [][]string {
	filas := len(matrizInterface)

	// Encontrar el número máximo de columnas
	maxColumnas := 0
	for i := 0; i < filas; i++ {
		if len(matrizInterface[i]) > maxColumnas {
			maxColumnas = len(matrizInterface[i])
		}
	}

	// Crear una nueva matriz de tipo int con las mismas filas y columnas máximas
	matrizInt := make([][]string, filas)
	for i := range matrizInt {
		matrizInt[i] = make([]string, maxColumnas)
	}

	// Recorrer la matriz de tipo interface{} y convertir cada elemento a int
	for i := 0; i < filas; i++ {
		for j := 0; j < len(matrizInterface[i]); j++ {
			valor, ok := matrizInterface[i][j].(string)
			if !ok {
				// Si el elemento no es un int, manejar el error como desees
				// En este caso, se establece a 0
				valor = ""
			}
			matrizInt[i][j] = valor
		}
	}

	return matrizInt
}
func ConvertirMatrizToInt(matrizInterface [][]interface{}) [][]int64 {
	filas := len(matrizInterface)

	// Encontrar el número máximo de columnas
	maxColumnas := 0
	for i := 0; i < filas; i++ {
		if len(matrizInterface[i]) > maxColumnas {
			maxColumnas = len(matrizInterface[i])
		}
	}

	// Crear una nueva matriz de tipo int con las mismas filas y columnas máximas
	matrizInt := make([][]int64, filas)
	for i := range matrizInt {
		matrizInt[i] = make([]int64, maxColumnas)
	}

	// Recorrer la matriz de tipo interface{} y convertir cada elemento a int
	for i := 0; i < filas; i++ {
		for j := 0; j < len(matrizInterface[i]); j++ {
			valor, ok := matrizInterface[i][j].(int64)
			if !ok {
				// Si el elemento no es un int, manejar el error como desees
				// En este caso, se establece a 0
				valor = 0
			}
			matrizInt[i][j] = valor
		}
	}

	return matrizInt
}
