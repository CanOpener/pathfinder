import React, { useEffect, useRef, useState } from 'react';
import axios from 'axios';

const SearchSpaceInfo = ({searchSpace, pollingInfo}) => {
  return (
    <div>
      {
        (() => {
          if (pollingInfo.status === "polling") {
            return (
              <div>
                <p>Polling ...</p>
                <p>Duration {pollingInfo.duration_ms / 1000}s</p>
              </div>
            )
          } else if (pollingInfo.status === "not_polling") {
            return (
              <div>
                <p>Generation Time: {searchSpace?.generation_date || ""}</p>
                <p>Generation Duration (ms): {searchSpace?.generation_duration_ms || ""}</p>
                <p>Node Plotting Algorithm: {searchSpace?.generation_job_parameters?.node_plotter_parameters?.node_plotter_id || ""}</p>
                <p>Node Connection Algorithm: {searchSpace?.generation_job_parameters?.node_connector_parameters?.node_connector_id || ""}</p>
                <p>Name Generator: {searchSpace?.generation_job_parameters?.name_generator_parameters?.name_generator_id || ""}</p>
                <p>Original Parameters: {JSON.stringify(searchSpace?.generation_job_parameters) || ""}</p>
              </div>
            )
          } else {
            return <div></div>;
          }
        })()
      }
    </div>
  );
}

export default SearchSpaceInfo;