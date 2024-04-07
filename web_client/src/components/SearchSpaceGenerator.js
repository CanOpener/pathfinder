import React, { useEffect, useRef, useState } from 'react';
import axios from 'axios';

const SearchSpaceGenerator = ({ searchSpace, setSearchSpace, pollingInfo, setPollingInfo, generationParameters, setGenerationParameters, dimensions, setSearchSpaces, setMode }) => {
  const handleParameterChange = (e) => {
    const { name, value, type, checked } = e.target;
    var newValue = value
    if (type === "number") {
      newValue = Number(value)
    }
    if (type === "checkbox") {
      newValue = checked
    }
    console.log(name, value)
    setGenerationParameters(prevParams => ({
      ...prevParams,
      [name]: newValue
    }));
  }
  
  const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const createJobParameters = {
        node_plotter_parameters: {
          node_plotter_id: generationParameters.node_plotter_id,
          node_count: generationParameters.node_count,
          minimum_distance: generationParameters.minimum_distance,
          maximum_plot_attempts: generationParameters.maximum_plot_attempts,
          grid_size_x: dimensions.width - 20,
          grid_size_y: dimensions.height - 20,
        },
        node_connector_parameters: {
          node_connector_id: generationParameters.node_connector_id,
          maximum_node_connection_count: generationParameters.maximum_node_connection_count,
        },
        name_generator_parameters: {
          name_generator_id: generationParameters.name_generator_id,
          allow_duplicates: generationParameters.allow_duplicates,
          maximum_sample_attempts: generationParameters.maximum_sample_attempts
        }
      }
      console.log(createJobParameters)
      const createJobResponse = await axios.post(`http://127.0.0.1:8080/generation_jobs`, createJobParameters, {
        headers: {
          'Content-Type': 'application/json'
        },
        timeout: 5000
      });
      if (!createJobResponse.data.success) {
        alert(`Error creating job: ${createJobResponse.data.error}`);
        return
      }
      
      let pollingStartTime = Date.now()
      let jobId = createJobResponse.data.generation_job_id
      var pollRequestCount = 0
      while (true) {
        await sleep(500);

        const pollResultsResponse = await axios.get(`http://127.0.0.1:8080/generation_jobs/${jobId}`)
        let currentTime = Date.now()

        if (pollResultsResponse.data.status.status === "error") {
          alert(`Error creating job: ${pollResultsResponse.data.status.error}`);
          setPollingInfo({status: "not_polling", duration_ms: (currentTime - pollingStartTime), request_count: pollRequestCount})
          return
        }

        if (pollResultsResponse.data.status.status === "finished") {
          setSearchSpace(pollResultsResponse.data.status.result)
          setPollingInfo({status: "not_polling", duration_ms: (currentTime - pollingStartTime), request_count: pollRequestCount})
          return
        }

        pollRequestCount += 1
        setPollingInfo({status: "polling", duration_ms: (currentTime - pollingStartTime), request_count: pollRequestCount})
      }
    } catch (error) {
      console.error("There was an error!", error);
    }
  };

  const handleLoadClick = async(e) => {
    try {
      const response = await axios.get('http://127.0.0.1:8080/search_spaces');
      setSearchSpaces(response.data.search_spaces);
      setMode("load")
    } catch (error) {
      console.error("Failed to fetch search spaces:", error);
    }
  };

  const handleSaveClick = async(e) => {
    try {
      const saveResponse = await axios.post(`http://127.0.0.1:8080/search_spaces`, searchSpace, {
        headers: {
          'Content-Type': 'application/json'
        },
        timeout: 5000
      });
      if (!saveResponse.data.success) {
        alert(`Error saving search state: ${saveResponse.data.error}`);
        return
      }
      var newSearchSpace = JSON.parse(JSON.stringify(searchSpace))
      newSearchSpace.name = ""
      newSearchSpace.nodes = {}
      setSearchSpace(newSearchSpace)
    } catch (error) {
      console.error("Failed to fetch search spaces:", error);
    }
  };

  const handleNameChange = (e) => {
    const { value, } = e.target;
    var newSearchSpace = JSON.parse(JSON.stringify(searchSpace))
    newSearchSpace.name = value
    setSearchSpace(newSearchSpace)
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <hr />
        <div className="button-row">
          <button type="submit">Generate</button>
          <button type="button"onClick={handleSaveClick}>Save</button>
          <button type="button" onClick={handleLoadClick}>Load</button>
        </div>

        <hr />
        {/* Saving parameters */}
        <label><b>Saving Parameters</b></label>
        <div>
          <label htmlFor="name">Name </label>
          <input type="text" name="name" value={searchSpace.name} onChange={handleNameChange} />
        </div>
        <hr />

        {/* Node plotter parameters */}
        <label><b>Node Plotter Parameters</b></label>

        <div>
          <label htmlFor="node_plotter_id">Node Plotter ID </label>
          <select name="node_plotter_id" value={generationParameters.node_plotter_id} onChange={handleParameterChange}>
            <option value="random">Random</option>
          </select>
        </div>

        <div>
          <label htmlFor="node_count">Node Count </label>
          <input type="number" name="node_count" value={generationParameters.node_count} onChange={handleParameterChange} />
        </div>

        <div>
          <label htmlFor="minimum_distance">Minimum Distance </label>
          <input type="number" name="minimum_distance" value={generationParameters.minimum_distance} onChange={handleParameterChange} />
        </div>

        <div>
          <label htmlFor="maximum_plot_attempts">Maximum Plot Attempts </label>
          <input type="number" name="maximum_plot_attempts" value={generationParameters.maximum_plot_attempts} onChange={handleParameterChange} />
        </div>

        <div>
          <label htmlFor="grid_size_x">Grid Size X </label>
          <input type="number" name="grid_size_x" value={dimensions?.width || 0} readOnly />
        </div>

        <div>
          <label htmlFor="grid_size_y">Grid Size Y </label>
          <input type="number" name="grid_size_y" value={dimensions?.height || 0} readOnly />
        </div>
        
        <hr />

        {/* Node connector parameters */}
        <label><b>Node connector parameters</b></label>

        <div>
          <label htmlFor="node_connector_id">Node Connector ID </label>
          <select name="node_connector_id" value={generationParameters.node_connector_id} onChange={handleParameterChange}>
            <option value="prim">Prim</option>
            <option value="maxn">Max N Connections</option>
            <option value="none">None</option>
          </select>
        </div>

        <div>
          <label htmlFor="maximum_node_connection_count">Maximum Node Connection Count </label>
          <input type="number" name="maximum_node_connection_count" value={generationParameters.maximum_node_connection_count} onChange={handleParameterChange} />
        </div>

        <hr />

        {/* Name generator parameters */}
        <label><b>Name generator parameters</b></label>

        <div>
          <label htmlFor="name_generator_id">Name Generator ID </label>
          <select name="name_generator_id" value={generationParameters.name_generator_id} onChange={handleParameterChange}>
            <option value="countries">Countries</option>
            <option value="cities">Cities</option>
            <option value="first_names">First Names</option>
            <option value="three_letters">Three Letters</option>
            <option value="uuid">UUID</option>
          </select>
        </div>

        <div>
          <label>
            <input type="checkbox" name="allow_duplicates" checked={generationParameters.allow_duplicates} onChange={handleParameterChange} />
            Allow Duplicates 
          </label>
        </div>

        <div>
          <label htmlFor="maximum_sample_attempts">Maximum Sample Attempts </label>
          <input type="number" name="maximum_sample_attempts" value={generationParameters.maximum_sample_attempts} onChange={handleParameterChange} />
        </div>

        <hr />
      </form>
    </div>
  );

}

export default SearchSpaceGenerator;