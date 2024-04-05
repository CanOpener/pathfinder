import React, { useEffect, useRef } from 'react';

const SearchSpaceDisplay = ({ searchSpace }) => {
  const canvasRef = useRef(null);

  useEffect(() => {
    if (searchSpace) {
      console.log(searchSpace)
      const canvas = canvasRef.current;
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height); // Clear canvas for redraw
      for (const nodeId in searchSpace.nodes) {
        const node = searchSpace.nodes[nodeId]
        ctx.beginPath();
        ctx.arc(node.x + 15, node.y + 15, 5, 0, 2 * Math.PI); // Draw a circle for each node
        ctx.fillText(nodeId, node.x + 10, node.y + 7); // Label each node
        ctx.stroke();

        // Draw lines to connected 
        for (const connectionId in node.connections) {
          const connectedNode = searchSpace.nodes[connectionId]
          if (connectedNode) {
            ctx.beginPath();
            ctx.moveTo(node.x + 15, node.y + 15); // Start line at current node
            ctx.lineTo(connectedNode.x + 15, connectedNode.y + 15); // Draw line to connected node
            ctx.stroke();
          }
        }
      }
    }
  }, [searchSpace]); // Redraw when searchSpace changes

  return (
    <div>
      <canvas ref={canvasRef} width={searchSpace.grid_size_x + 30} height={searchSpace.grid_size_y + 30} style={{ border: "1px solid #000" }}></canvas>
      <div>
        <p>Generation Time: {searchSpace?.generation_date || ""}</p>
        <p>Generation Duration (ms): {searchSpace?.generation_duration_ms || ""}</p>
        <p>Node Plotting Algorithm: {searchSpace?.generation_job_parameters?.node_plotter_parameters?.node_plotter_id || ""}</p>
        <p>Node Connection Algorithm: {searchSpace?.generation_job_parameters?.node_connector_parameters?.node_connector_id || ""}</p>
        <p>Name Generator: {searchSpace?.generation_job_parameters?.name_generator_parameters?.name_generator_id || ""}</p>
        <p>Original Parameters: {JSON.stringify(searchSpace?.generation_job_parameters) || ""}</p>
      </div>
    </div>
  );
}

export default SearchSpaceDisplay;