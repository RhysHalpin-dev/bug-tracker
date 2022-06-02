import React from 'react';
import styled from 'styled-components';

const DashTickets: React.FC = () => {
  return <Container>TICKETS</Container>;
};

export default DashTickets;

const Container = styled.div`
  border: solid 3px blue;
  width: 90%;
  display: flex;
  flex-direction: row;
  align-items: stretch;
`;
