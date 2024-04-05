import React, { useEffect, useRef, useState } from 'react';
import axios from 'axios';

const SearchSpaceGenerator = ({ searchSpace, setSearchSpace }) => {
  const [generationParameters, setGenerationParameters] = useState({
    node_plotter_id: "random",
    node_count: 500,
    minimum_distance: 25,
    maximum_plot_attempts: 10000,
    grid_size_x: 2300,
    grid_size_y: 900,
    node_connector_id: "prim",
    maximum_node_connection_count: 0,
    name_generator_id: "countries",
    allow_duplicates: true,
    maximum_sample_attempts: 0
  })

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
          grid_size_x: generationParameters.grid_size_x,
          grid_size_y: generationParameters.grid_size_y,
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

      let jobId = createJobResponse.data.generation_job_id
      while (true) {
        await sleep(500);

        const pollResultsResponse = await axios.get(`http://127.0.0.1:8080/generation_jobs/${jobId}`)
        if (pollResultsResponse.data.status.status === "error") {
          alert(`Error creating job: ${pollResultsResponse.data.status.error}`);
          return
        }

        if (pollResultsResponse.data.status.status === "finished") {
          setSearchSpace(pollResultsResponse.data.status.result)
          return
        }
      }
    } catch (error) {
      console.error("There was an error!", error);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        {/* Node plotter parameters */}
        <select name="node_plotter_id" value={generationParameters.node_plotter_id} onChange={handleParameterChange}>
          <option value="random">Random</option>
        </select>
        <input type="number" name="node_count" value={generationParameters.node_count} onChange={handleParameterChange} />
        <input type="number" name="minimum_distance" value={generationParameters.minimum_distance} onChange={handleParameterChange} />
        <input type="number" name="maximum_plot_attempts" value={generationParameters.maximum_plot_attempts} onChange={handleParameterChange} />
        <input type="number" name="grid_size_x" value={generationParameters.grid_size_x} onChange={handleParameterChange} />
        <input type="number" name="grid_size_y" value={generationParameters.grid_size_y} onChange={handleParameterChange} />

        {/* Node connector parameters */}
        <select name="node_connector_id" value={generationParameters.node_connector_id} onChange={handleParameterChange}>
          <option value="prim">Prim</option>
          <option value="maxn">Max N Connections</option>
          <option value="none">None</option>
        </select>
        <input type="number" name="maximum_node_connection_count" value={generationParameters.maximum_node_connection_count} onChange={handleParameterChange} />

        {/* Name generator parameters */}
        <select name="name_generator_id" value={generationParameters.name_generator_id} onChange={handleParameterChange}>
          <option value="countries">Countries</option>
          <option value="cities">Cities</option>
          <option value="first_names">First Names</option>
          <option value="three_letters">Three Letters</option>
        </select>
        <input type="checkbox" name="allow_duplicates" checked={generationParameters.allow_duplicates} onChange={handleParameterChange} />
        <input type="number" name="maximum_sample_attempts" value={generationParameters.maximum_sample_attempts} onChange={handleParameterChange} />

        <button type="submit">Generate</button>
      </form>
    </div>
  );

}

export default SearchSpaceGenerator;