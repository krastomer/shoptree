import React, { useState, useRef, useEffect } from "react";
import { Stage, Layer, Rect, Image as KonvaImag } from "react-konva";
import "./index.css";
import useImage from "use-image";
import A from "../../asset/all_product/Product_A.png";
import B from "../../asset/all_product/Product_B.png";
import C from "../../asset/all_product/Product_C.png";
import D from "../../asset/all_product/Product_D.png";
import E from "../../asset/all_product/Product_E.png";
import F from "../../asset/all_product/Product_F.png";
const WIDTH = 1000;
const HEIGHT = 1000;
const endY = Math.abs(window.innerHeight * 2);
const endX = Math.abs(window.innerWidth * 2);
const grid = [
  [A, "yellow"],
  ["green", "blue"],
];
const Canvas = (x, y) => {
  const canvasRef = useRef(null);
  useEffect(() => {
    const canvas = canvasRef.current;
    const context = canvas.getContext("2d");
    var obj1 = new Image();
    obj1.src = A;
    obj1.onload = function () {
      context.drawImage(this, 500, 0);
    };
  }, []);

  return <canvas ref={canvasRef} width={endX} height={endY} xmlns="http://www.w3.org/1999/xhtml" />;
};
export default function Home() {
  const [stagePos, setStagePos] = React.useState({ x: 0, y: 0 });
  const startX = 0;
 
  const startY = 0;
  
  const gridComponents = [];
  var i = 0;

  for (let x = startX; x < endX; x += WIDTH) {
    for (let y = startY; y < endY; y += HEIGHT) {
      if (i === 4) {
        i = 0;
      }
      // gridComponents.push(<Canvas x ={x} y = {y}></Canvas>)
    }
  }
  return (
    <div>
      <Canvas></Canvas>
    </div>
    // <Stage
    //   x={-(window.innerWidth * 2) / 2}
    //   y={-(window.innerHeight * 2) / 2}
    //   width={window.innerWidth}
    //   height={window.innerHeight}
    //   draggable
    //   onDragEnd={(e) => {
    //     setStagePos(e.currentTarget.position());
    //     console.log(e.currentTarget.position().x);
    //   }}
    // >
    //   <Layer>{gridComponents}</Layer>
    // </Stage>
  );
}
