import React, { useState } from 'react';

function Image() {
  const [svgContent, setSvgContent] = useState(null);

  // Función para cargar el archivo SVG
  const loadSvg = async () => {
    try {
      const response = await fetch(' http://127.0.0.1:8080/graph');
      const data = await response.text();
      setSvgContent(data);
    } catch (error) {
      console.error('Error al cargar el archivo SVG:', error);
    }
  };
 
  return (
    <div>
      <button style={buttonStyle} onClick={loadSvg}>Cargar SVG</button>
      <h3 style={{textAlign:"center" ,color:"white",fontSize:"25px"}}>CST</h3>
      {svgContent && (
        <div style={svgContainerStyle}>
          <div
            dangerouslySetInnerHTML={{ __html: svgContent }}
            style={svgStyle}
          />
        </div>
      )}
    </div>
  );

}
const buttonStyle = {
    marginTop: '20px', 
    marginBottom: '20px',
  };
  const svgContainerStyle = {
    border: '8px solid black', // Agrega el borde aquí
    padding: '10px', // Añade un poco de espacio interior para separar el borde del contenido
    width: 'auto', 
    height: 'auto', 
    margin: "0 auto",
    borderRadius:"15px",
    backgroundColor:"white"
    
  };
  const svgStyle = {
    width: '100%', // Establece el ancho de la imagen al 100% del contenedor
  };
export default Image;
