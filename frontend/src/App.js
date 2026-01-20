import React, { useState } from 'react';

function App() {
  const [ticketNumber, setTicketNumber] = useState('');
  const [seatPosition, setSeatPosition] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`http://localhost:8080/seat?number=${ticketNumber}`);
      const data = await response.json();
      setSeatPosition(data.position);
    } catch (error) {
      console.error('Error fetching seat position:', error);
      setSeatPosition('Error fetching position');
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <h1>Railway Seat Allocation</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Ticket Number:
          <input
            type="number"
            value={ticketNumber}
            onChange={(e) => setTicketNumber(e.target.value)}
            required
          />
        </label>
        <button type="submit">Get Seat Position</button>
      </form>
      {seatPosition && <p>Your seat position: {seatPosition}</p>}
    </div>
  );
}

export default App;