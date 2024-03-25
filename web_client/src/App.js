import React, { useState } from 'react';
import axios from 'axios';

function App() {
  const [parameters, setParameters] = useState({
    node_count: '25',
    name_generator: 'default',
    algorithm: 'default'
  });

  const [searchSpace, setSearchSpace] = useState(null); // Define searchSpace and its setter here


  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setParameters(prevParams => ({
      ...prevParams,
      [name]: value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.get(`http://127.0.0.1:8080/generate_search_space`, { params: parameters });
      console.log(response);
      setSearchSpace(response.data);
    } catch (error) {
      console.error("There was an error!", error);
    }
  };

  // Add this useEffect hook inside your App component, but outside the return statement
  React.useEffect(() => {
    if (searchSpace) {
      const canvas = document.getElementById('searchSpaceCanvas');
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height); // Clear canvas for redraw
      Object.values(searchSpace.search_space.nodes).forEach(node => {
        ctx.beginPath();
        ctx.arc(node.x, node.y, 5, 0, 2 * Math.PI); // Draw a circle for each node
        ctx.fillText(node.id, node.x + 10, node.y + 5); // Label each node
        ctx.stroke();
      });
    }
  }, [searchSpace]); // Redraw when searchSpace changes

  return (
    <div className="App">
      <form onSubmit={handleSubmit}>
        <input
          type="number"
          name="node_count"
          value={parameters.node_count}
          onChange={handleInputChange} />
        <select name="name_generator" value={parameters.name_generator} onChange={handleInputChange}>
          <option value="default">Default</option>
          {/* Add other name_generator options here */}
        </select>
        <select name="algorithm" value={parameters.algorithm} onChange={handleInputChange}>
          <option value="default">Default</option>
          {/* Add other algorithm options here */}
        </select>
        <button type="submit">Generate</button>
      </form>
      {
        searchSpace && (
          <div>
            <canvas id="searchSpaceCanvas" width="1000" height="1000" style={{ border: "1px solid #000" }}></canvas>
            <div>
              <p>ID: {searchSpace.search_space.id}</p>
              <p>Generation Time: {searchSpace.search_space.generation_date}</p>
              <p>Algorithm: {searchSpace.search_space.algorithm}</p>
              {/* Display other meta information similarly */}
            </div>
          </div>
        )
      }
    </div>
  );
}

export default App;
