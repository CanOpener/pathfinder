import React, { useEffect, useRef } from 'react';

const SearchSpaceDisplay = ({ searchSpace, pollingInfo, dimensions }) => {
  const canvasRef = useRef(null);
  const gridSectionRef = useRef(null);

  // Effect to adjust canvas size whenever dimensions prop changes
  useEffect(() => {
    const canvas = canvasRef.current;
    if (canvas && dimensions) {
      canvas.width = dimensions.width;
      canvas.height = dimensions.height;

      // Redrawing logic here if needed, or it could stay in the separate effect for searchSpace updates
    }
  }, [dimensions]); // React to changes in dimensions

  useEffect(() => {
    if (searchSpace && (pollingInfo.status === "not_polling")) {
      console.log(searchSpace)
      const canvas = canvasRef.current;
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height); // Clear canvas for redraw
      for (const nodeId in searchSpace.nodes) {
        const node = searchSpace.nodes[nodeId]
        ctx.beginPath();
        ctx.arc(node.x + 10, node.y + 10, 5, 0, 2 * Math.PI); // Draw a circle for each node
        ctx.fillText(nodeId, node.x + 10, node.y + 7); // Label each node
        ctx.stroke();

        // Draw lines to connected 
        for (const connectionId in node.connections) {
          const connectedNode = searchSpace.nodes[connectionId]
          if (connectedNode) {
            ctx.beginPath();
            ctx.moveTo(node.x + 10, node.y + 10); // Start line at current node
            ctx.lineTo(connectedNode.x + 10, connectedNode.y + 10); // Draw line to connected node
            ctx.stroke();
          }
        }
      }
    }
  }, [searchSpace]); // Redraw when searchSpace changes

  return (
    <div ref={gridSectionRef}>
      <canvas ref={canvasRef} style={{ border: "1px solid #000" }}></canvas>
    </div>
  );
}

export default SearchSpaceDisplay;