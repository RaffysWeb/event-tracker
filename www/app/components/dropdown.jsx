import React from 'react';

const Dropdown = ({ options, onChange, selectedAction }) => {

  const handleSelectChange = (e) => {
    const value = e.target.value;
    onChange(value);
  };

  return (
    <select
      className="border rounded p-2 mb-1"
      value={selectedAction}
      onChange={handleSelectChange}
    >
      <option value="">Select...</option>
      {options.map((option) => (
        <option key={option.value} value={option.value}>
          {option.label}
        </option>
      ))}
    </select>
  );
};

export default Dropdown;
