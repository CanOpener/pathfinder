import React, { useState, useEffect, useRef } from 'react';
import SearchSpaceDisplay from './components/SearchSpaceDisplay'
import SearchSpaceGenerator from './components/SearchSpaceGenerator'
import SearchSpaceInfo from './components/SearchSpaceInfo'
import SearchSpaceSearcher from './components/SearchSpaceSearcher'
import './App.css';

function App() {
  const searchSpaceDisplayRef = useRef(null);

  const [mode, setMode] = useState("generation")

  const [searchSpace, setSearchSpace] = useState({
    name: "",
    generation_date: "",
    generation_duration_ms: 0,
    generation_job_parameters: {},
    grid_size_x: 2300,
    grid_size_y: 900,
    nodes: {}
  });

  const [pollingInfo, setPollingInfo] = useState({
    status: "not_polling",
    duration_ms: 0,
  })

  const [generationParameters, setGenerationParameters] = useState({
    node_plotter_id: "random",
    node_count: 500,
    minimum_distance: 25,
    maximum_plot_attempts: 10000,
    node_connector_id: "prim",
    maximum_node_connection_count: 0,
    name_generator_id: "countries",
    allow_duplicates: true,
    maximum_sample_attempts: 0
  })

  const [dimensions, setDimensions] = useState(null);

  useEffect(() => {
    if (searchSpaceDisplayRef.current) {
      let width = searchSpaceDisplayRef.current.offsetWidth - 10
      let height = searchSpaceDisplayRef.current.offsetHeight
      const updateDimensions = () => {
        setDimensions({
          width: width,
          height: height,
        });
      };
  
      updateDimensions();
    }
  }, []);

  return (
    <div className="App">
      <div className="left-section">
        <div className="search-space-display" ref={searchSpaceDisplayRef}><SearchSpaceDisplay searchSpace={searchSpace} 
                                                                                              pollingInfo={pollingInfo}
                                                                                              dimensions={dimensions}
                                                                                              setDimensions={setDimensions}/></div>
        <div className="search-space-info"><SearchSpaceInfo searchSpace={searchSpace} pollingInfo={pollingInfo} /></div>
      </div>
      <div className="right-section">
        {
          (() => {
            if (mode === "generation") {
              return <SearchSpaceGenerator searchSpace={searchSpace} 
                                           setSearchSpace={setSearchSpace}
                                           pollingInfo={pollingInfo}
                                           setPollingInfo={setPollingInfo}
                                           generationParameters={generationParameters}
                                           setGenerationParameters={setGenerationParameters}
                                           dimensions={dimensions} />;
            } else if (mode === "search") {
              return <SearchSpaceSearcher />;
            } else {
              return <div></div>;
            }
          })()
        }
      </div>
    </div>
  );
}

export default App;
