import React, { useState } from "react";
import { Stage, Layer, Rect, Image } from "react-konva";
import "./index.css";
import useImage from "use-image";
import A from "../../asset/all_product/Product_A.png";
import B from "../../asset/all_product/Product_B.png";
import C from "../../asset/all_product/Product_C.png";
import D from "../../asset/all_product/Product_D.png";
import E from "../../asset/all_product/Product_E.png";
import F from "../../asset/all_product/Product_F.png";
const WIDTH = 100;
const HEIGHT = 100;

const grid = [
  [A, "yellow"],
  ["green", "blue"]
];


export default function Home() {
  const [image] = useImage(A)
  const [stagePos, setStagePos] = React.useState({ x: 0, y: 0 });
  const startX = 0
  const endX = Math.abs(window.innerWidth * 2)
  const startY = 0
  const endY =Math.abs(window.innerHeight*2)
  const gridComponents = [];
  var i = 0;
  let mod = 0;
  for (var x = startX; x < endX; x += WIDTH) {
    for (var y = startY; y < endY; y += HEIGHT) {
      if (i === 4) {
        i = 0;
      }
      // if(x <100&&y <100){
      //   gridComponents.push(
      //     <Image image = {image} x = {x} y = {y}></Image>
      //   );
      // }
      // console.log(gridComponents)
      const indexX = Math.abs(x / WIDTH) % grid.length;
      const indexY = Math.abs(y / HEIGHT) % grid[0].length;
      gridComponents.push(
        <Rect
          x={x}
          y={y}
          width={WIDTH}
          height={HEIGHT}
          fill={grid[indexX][indexY]}
          stroke="black"
        />
      );
    }
  }
  return (
      <Stage
        x={-(window.innerWidth * 2)/2}
        y={-(window.innerHeight*2)/2}
        width={window.innerWidth}
        height={window.innerHeight}
        draggable
        onDragEnd={(e) => {
          setStagePos(e.currentTarget.position());
          console.log(e.currentTarget.position().x)
        }}
      >
        <Layer>{gridComponents}</Layer>
      </Stage>
  );
}
