import React, { useState, useRef } from 'react';
import AceEditor from 'react-ace';
import 'ace-builds/src-noconflict/theme-xcode'; // Importa el tema que prefieras
import 'ace-builds/src-noconflict/mode-swift';
import './App.css'

function App() {
  const [editorContent, setEditorContent] = useState('');
  const [executionResult, setExecutionResult] = useState('Resultado de la ejecución...');
  const fileInputRef = useRef(null);

  const handleFileButtonClick = () => {
    fileInputRef.current.click();
  };

  const handleFileInputChange = async (e) => {
    const file = e.target.files[0];

    if (file) {
      const fileContent = await readFileAsync(file);
      setEditorContent(fileContent);
    }
  };

  const handleSaveButtonClick = () => {
    const blob = new Blob([editorContent], { type: 'text/plain' });
    const a = document.createElement('a');
    a.href = URL.createObjectURL(blob);
    a.download = 't_swift_doc.txt';
    a.click();
  };

  const handleRunButtonClick = async () => {
    const res = await fetch('http://localhost:8080/analyze', {
      method: 'POST',
      headers: {
        'Content-Type': 'text/plain',
      },
      body: editorContent,
    });
    const response = await res.json(); // Recibir la respuesta como JSON
   // Acceder al contenido utilizando la propiedad "data"

  setExecutionResult(response.prints);
  console.log(executionResult);
  };

  const readFileAsync = (file) => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();

      reader.onload = (event) => {
        resolve(event.target.result);
      };

      reader.onerror = (error) => {
        reject(error);
      };

      reader.readAsText(file);
    });
  };

  return (
    <>
      <h1>T-Swift</h1>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <div style={{ width: '60%', marginLeft: "40px" }}>

          <AceEditor
            mode="swift"
            theme="xcode"
            name="code-editor"
            value={editorContent}
            onChange={(newContent) => setEditorContent(newContent)}
            editorProps={{ $blockScrolling: true }}

            style={{ width: '100%', height: '300px', borderRadius: '5px',fontSize: '15px'}}
          />
        </div>
        <div style={{ marginLeft: '20px', width: "30%" }}>
          <div style={{ marginBottom: '20px' }}>

            <button onClick={handleFileButtonClick} style={{ width: '45%', marginRight: "38px" }}>Abrir</button>
            <input
              type="file"
              ref={fileInputRef}
              accept=".txt"
              style={{ display: 'none' }}
              onChange={handleFileInputChange}
            />
            <button onClick={handleSaveButtonClick} style={{ width: '45%' }}>Guardar</button>

          </div>
          <div style={{ marginBottom: '20px' }}>
            <button style={{ width: '45%', marginRight: "38px" }}>Guardar como</button>
            <button style={{ width: '45%' }}>Mostrar errores</button>
          </div>
          <div>
            <button onClick={handleRunButtonClick}
              style={{ width: '100%' }}>Ejecutar</button>
          </div>
        </div>
      </div>
      <textarea
        value={executionResult}
        readOnly
        style={{ width: '90%', height: '150px', marginLeft: "40px",
                marginTop: '20px', border: '1px solid #ccc', borderRadius: '5px', 
                padding: '10px', overflow: 'auto', fontSize: '15px'}}
      />
    </>

  );

}
const buttonStyle = {
  fontSize: '16px',
  padding: '10px 20px',
  border: '1px solid #ccc',
  borderRadius: '5px',
};
export default App
