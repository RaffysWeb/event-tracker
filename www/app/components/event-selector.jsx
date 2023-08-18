"use client"

import React, { useState } from 'react';
import Dropdown from './dropdown';

const EventSelector = () => {
  const users = [
    { value: 'user1', label: 'User 1' },
    { value: 'user2', label: 'User 2' },
    // Add more users...
  ];

  const actionTypes = [
    { value: 'action1', label: 'Action 1' },
    { value: 'action2', label: 'Action 2' },
    // Add more actions...
  ];

  const [selectedUser, setSelectedUser] = useState('');
  const [selectedAction, setSelectedAction] = useState('');

  const handleUserChange = (user) => {
    setSelectedUser(user);
    setSelectedAction('');
  };

  const handleActionChange = (action) => {
    setSelectedAction(action);
  };

  return (
    <div className="container mx-auto p-8">
      <h1 className="text-2xl font-semibold mb-4">Event Metrics</h1>
      <div className='flex flex-col max-w-xs'>
        <Dropdown options={users} onChange={handleUserChange} />
        {selectedUser && (
          <Dropdown options={actionTypes} selectedAction={selectedAction} onChange={handleActionChange} />
        )}
        {selectedAction && (
          <div className="mt-4">
            {/* Display metrics for the selected user and action */}
            Metrics for {selectedUser} performing {selectedAction}
          </div>
        )}
      </div>
    </div>
  );
};

export default EventSelector;
