import React, { useEffect, useState } from 'react';
import axios from 'axios';

const SearchSpaceLoader = ({ searchSpaces, setSearchSpace, setMode }) => {
  // Function to handle clicking on a search space
  const handleClick = async(e, id) => {
    try {
      const showRespoonse = await axios.get(`http://127.0.0.1:8080/search_spaces/${id}`);
      if (!showRespoonse.data.success) {
        alert(`Error viewing search state: ${showRespoonse.data.error}`);
        return
      }
      setSearchSpace(showRespoonse.data.search_space)
      setMode("search")
    } catch (error) {
      alert(`Error viewing search state: ${error}`);
    }
  };
  
  const handleCancelClick = async(e) => {
    setMode("generation")
  };

  return (
    <div>
      <hr />
      <div className="button-row">
        <button type="button"onClick={handleCancelClick}>Cancel</button>
      </div>
      <hr />

      <div>
          {searchSpaces.map((space, index) => (
            <div class="vertical-labels" onClick={(e) => handleClick(e, space.id)} key={index}>
              <label htmlFor="name"><b>{space.name}</b></label>
              <label htmlFor="generation_date">Generation Date : <i>{space.generation_date}</i></label>
              <label htmlFor="node_count">Node Count : <i>{space.node_count}</i></label>
              <label htmlFor="grid_size">Node Count : <i>{space.grid_size_x}x{space.grid_size_y}</i></label>
              <hr />
            </div>
          ))}
      </div>
    </div>
  );
};

export default SearchSpaceLoader;
