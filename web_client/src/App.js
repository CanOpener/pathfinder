import React, { useState } from 'react';
import SearchSpaceDisplay from './components/SearchSpaceDisplay'
import SearchSpaceGenerator from './components/SearchSpaceGenerator'

function App() {
  const [searchSpace, setSearchSpace] = useState({
    name: "",
    generation_date: "",
    generation_duration_ms: 0,
    generation_job_parameters: {},
    grid_size_x: 2300,
    grid_size_y: 900,
    nodes: {}
  });

  return (
    <div className="App">
      {
        searchSpace && <SearchSpaceDisplay searchSpace={searchSpace} />
      }
      {
        searchSpace && <SearchSpaceGenerator searchSpace={searchSpace} setSearchSpace={setSearchSpace} />
      }
    </div>
  );
}

export default App;
