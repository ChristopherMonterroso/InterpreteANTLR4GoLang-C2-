package parsermatrix

func ConvertirMatriz3ToString(matrizInterface [][][]interface{}) [][][]string {
	capas := len(matrizInterface)

	// Encontrar el número máximo de filas y columnas
	maxFilas := 0
	maxColumnas := 0
	for i := 0; i < capas; i++ {
		if len(matrizInterface[i]) > maxFilas {
			maxFilas = len(matrizInterface[i])
		}
		for j := 0; j < len(matrizInterface[i]); j++ {
			if len(matrizInterface[i][j]) > maxColumnas {
				maxColumnas = len(matrizInterface[i][j])
			}
		}
	}

	// Crear una nueva matriz de tipo int con las mismas capas, filas y columnas máximas
	matrizInt := make([][][]string, capas)
	for i := range matrizInt {
		matrizInt[i] = make([][]string, maxFilas)
		for j := range matrizInt[i] {
			matrizInt[i][j] = make([]string, maxColumnas)
		}
	}

	// Recorrer la matriz de tipo interface{} y convertir cada elemento a int
	for i := 0; i < capas; i++ {
		for j := 0; j < len(matrizInterface[i]); j++ {
			for k := 0; k < len(matrizInterface[i][j]); k++ {
				valor, ok := matrizInterface[i][j][k].(string)
				if !ok {
					// Si el elemento no es un int, manejar el error como desees
					// En este caso, se establece a 0
					valor = ""
				}
				matrizInt[i][j][k] = valor
			}
		}
	}

	return matrizInt
}
func ConvertirMatriz3ToInt(matrizInterface [][][]interface{}) [][][]int64 {
	capas := len(matrizInterface)

	// Encontrar el número máximo de filas y columnas
	maxFilas := 0
	maxColumnas := 0
	for i := 0; i < capas; i++ {
		if len(matrizInterface[i]) > maxFilas {
			maxFilas = len(matrizInterface[i])
		}
		for j := 0; j < len(matrizInterface[i]); j++ {
			if len(matrizInterface[i][j]) > maxColumnas {
				maxColumnas = len(matrizInterface[i][j])
			}
		}
	}

	// Crear una nueva matriz de tipo int con las mismas capas, filas y columnas máximas
	matrizInt := make([][][]int64, capas)
	for i := range matrizInt {
		matrizInt[i] = make([][]int64, maxFilas)
		for j := range matrizInt[i] {
			matrizInt[i][j] = make([]int64, maxColumnas)
		}
	}

	// Recorrer la matriz de tipo interface{} y convertir cada elemento a int
	for i := 0; i < capas; i++ {
		for j := 0; j < len(matrizInterface[i]); j++ {
			for k := 0; k < len(matrizInterface[i][j]); k++ {
				valor, ok := matrizInterface[i][j][k].(int64)
				if !ok {
					// Si el elemento no es un int, manejar el error como desees
					// En este caso, se establece a 0
					valor = 0
				}
				matrizInt[i][j][k] = valor
			}
		}
	}

	return matrizInt
}
