import React from 'react';
import { Box } from '../styles/box';
import Chart, { Props } from 'react-apexcharts';

const state: Props['series'] = [
   {
      name: 'Productivity',
      data: [
         {
            x: 'Monday',
            y: 6,
         },
         {
            x: 'Tuesday',
            y: 5.5,
         },
         {
            x: 'Wednesday',
            y: 3,
         },
         {
            x: 'Thursday',
            y: 2,
         },
         {
            x: 'Friday',
            y: 6.6,
         },
         {
            x: 'Saturday',
            y: 1,
         }
      ]
   },
   {
      name: 'Non productive',
      data: [
         {
            x: 'Monday',
            y: 4,
         },
         {
            x: 'Tuesday',
            y: 3.5,
         },
         {
            x: 'Wednesday',
            y: 2,
         },
         {
            x: 'Thursday',
            y: 1,
         },
         {
            x: 'Friday',
            y: 5,
         },
         {
            x: 'Saturday',
            y: 0.5,
         }
      ]
   }
];

const options: Props['options'] = {
   chart: {
      type: 'bar',
      animations: {
         easing: 'linear',
         speed: 300,
      },
      sparkline: {
         enabled: false,
      },
      brush: {
         enabled: false,
      },
      id: 'basic-bar',
      fontFamily: 'Inter, sans-serif',
      foreColor: 'var(--nextui-colors-accents9)',
      stacked: false,
      toolbar: {
         show: false,
      },
   },
   legend: {
      show: true,
      position: 'top',
   },
   xaxis: {
      categories: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
      labels: {
         // show: false,
         style: {
            colors: 'var(--nextui-colors-accents8)',
            fontFamily: 'Inter, sans-serif',
         },
      },
      axisBorder: {
         color: 'var(--nextui-colors-border)',
      },
      axisTicks: {
         color: 'var(--nextui-colors-border)',
      },
   },
   yaxis: {
      labels: {
         formatter:(hours)=>{return hours + 'hrs'},
         style: {
            colors: 'var(--nextui-colors-accents8)',
            fontFamily: 'Inter, sans-serif',
         },       
      },
      title: {
         text: 'Hours',
         style: {
            color: 'var(--nextui-colors-accents8)',
            fontFamily: 'Inter, sans-serif',
         },
      }
   },
   tooltip: {
      enabled: false,
   },
   grid: {
      show: true,
      borderColor: 'var(--nextui-colors-border)',
      strokeDashArray: 0,
      position: 'back',
   },
   stroke: {
      curve: 'smooth',
      fill: {
         colors: ['red'],
      },
   },
   // @ts-ignore
   markers: false,
   dataLabels: {
      enabled: false,
   }
};

export const Steam = () => {
   return (
      <>
         <Box
            css={{
               width: '100%',
               zIndex: 5,
            }}
         >
            <div id="chart">
               <Chart
                  options={options}
                  series={state}
                  type="bar"
                  height={425}
               />
            </div>
         </Box>
      </>
   );
};
