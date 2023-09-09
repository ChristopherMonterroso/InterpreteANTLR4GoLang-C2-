import React from 'react';

const TableError = ({ title, arrayErrors }) => {
  return (
    <div>
      <h2>{title}</h2>
      {arrayErrors != null ? (
        <table id='lexersTable'>
          <thead>
            <tr>
              <th>Mensaje</th>
              <th>Línea</th>
              <th>Columna</th>
            </tr>
          </thead>
          <tbody>
            {arrayErrors.map((error, index) => (
              <tr key={index}>
                <td>{error.ErrorMessage}</td>
                <td>{error.Line}</td>
                <td>{error.Column}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p style={{ color: "white" }}>No se encontraron {title}.</p>
      )}
    </div>
  );
}

export default TableError;