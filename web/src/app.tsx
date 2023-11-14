import React, { useEffect } from "react";
import Log from "./utils/log";

//Layers

const App: React.FC = () => {

  useEffect(() => {
    Log.identify('Test');
  }, [])


  return (
    <>
      <h1>Hello World!</h1>
    </>
  );
};

export default App;
