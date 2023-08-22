"use client"
import React from 'react';
import EventSelector from './components/event-selector';
import Chart from './components/chart';

const Home = () => {
  return (
    <div>
      <EventSelector />
      <Chart />
    </div>
  );
};

export default Home;
