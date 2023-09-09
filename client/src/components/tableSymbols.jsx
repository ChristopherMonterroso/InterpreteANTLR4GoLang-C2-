import React from 'react';

const TableSymbol = ({ array }) => {
  // Verifica si array es null o undefined
  if (!array) {
    return (<div>
      <h2>Tabla de símbolos:</h2>
      <p style={{ color: "white" }}>No se encontraron símbolos.</p>;
    </div>
    
    )
  }

  // Eliminar duplicados por atributos exactamente iguales
  const uniqueArray = Array.from(new Set(array.map(JSON.stringify))).map(JSON.parse);

  return (
    <div>
      <h2>Tabla de símbolos:</h2>
      {uniqueArray.length > 0 ? (
        <table id='lexersTable'>
          <thead>
            <tr>
              <th>ID</th>
              <th>Tipo</th>
              <th>Tipo símbolo</th>
              
              <th>Ámbito</th>
              <th>Línea</th>
              <th>Columna</th>
            </tr>
          </thead>
          <tbody>
            {uniqueArray.map((symbol, index) => (
              <tr key={index}>
                <td>{symbol.Id}</td>
                <td>{symbol.TipoSymbol}</td>
                <td>{symbol.Tipo}</td>
                <td>{symbol.Ambit}</td>
                <td>{symbol.Line}</td>
                <td>{symbol.Column}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p style={{ color: "white" }}>No se encontraron símbolos.</p>
      )}
    </div>
  );
}

export default TableSymbol;
