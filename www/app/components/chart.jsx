import React from 'react';
import {
  Chart as ChartJS,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  TimeScale
} from 'chart.js';
import { Line } from 'react-chartjs-2';
import useWebSocket from 'react-use-websocket';
import 'chartjs-adapter-date-fns';
import { enUS } from 'date-fns/locale';

ChartJS.register(LinearScale, PointElement, LineElement, Title, Tooltip, Legend, TimeScale);

function App() {
  const colors = [
    'rgba(255, 99, 132, 1)',
    'rgba(255, 159, 64, 1)',
    'rgba(255, 205, 86, 1)',
    'rgba(75, 192, 192, 1)',
  ]
  const { lastMessage } = useWebSocket("ws://localhost:4001/live-events");
  const data = lastMessage && JSON.parse(lastMessage?.data) || []

  console.log(data)
  const eventTypes = [...new Set(data.map(item => item.event_type))];

  const datasets = eventTypes.map((eventType, index) => {
    const filteredData = data.filter(item => item.event_type === eventType);

    return {
      label: capitalize(eventType),
      data: filteredData.map(item => ({ x: item._time, y: item._value })),
      fill: false,
      borderColor: colors[index],
      tension: 0.1,
    };
  });

  const chartData = {
    datasets,
  };

  console.log(data[0]?._time)

  return (
    <div className="App" >
      <Line
        data={chartData}
        options={{
          scales: {
            x: {
              type: 'time',
              min: data[0]?._time,
              max: data[data.length - 1]?._time,
              adapters: {
                date: {
                  locale: enUS,
                },
              },
              ticks: {
                stepSize: 10,
              },
              time: {
                unit: 'second',
              },
              title: {
                display: true,
                text: 'Time',
              },
            },
            y: {
              beginAtZero: false,
              title: {
                display: true,
                text: 'Value',
              },
            },
          },
        }}
      />
    </div >
  );

  function capitalize(str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
  }
}

export default App;
