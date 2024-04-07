import React, { useEffect, useRef, useState } from 'react';
import axios from 'axios';

const SearchSpaceSearcher = ({ searchSpace, setSearchSpace, setSearchSpaces, setMode }) => {

  const handleLoadClick = async(e) => {
    try {
      const response = await axios.get('http://127.0.0.1:8080/search_spaces');
      setSearchSpaces(response.data.search_spaces);
      setMode("load")
    } catch (error) {
      console.error("Failed to fetch search spaces:", error);
    }
  };

  const handleGenerateClick = (e) => {
    var newSearchSpace = JSON.parse(JSON.stringify(searchSpace))
    newSearchSpace.name = ""
    newSearchSpace.nodes = {}
    setSearchSpace(newSearchSpace)
    setMode("generation")
  }

  return (
    <div>
      <hr />
      <div className="button-row">
        <button type="button"onClick={handleGenerateClick}>Generation</button>
        <button type="button" onClick={handleLoadClick}>Load</button>
      </div>
      <hr />
    </div>
  );
}

export default SearchSpaceSearcher;