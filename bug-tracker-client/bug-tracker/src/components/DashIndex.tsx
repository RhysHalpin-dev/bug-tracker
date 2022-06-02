import React from 'react';
import styled from 'styled-components';
import Feed from './Feed';
import Profile from './Profile';

const DashIndex: React.FC = () => {
  return (
    <Container>
      <Feed />
      <Profile />
    </Container>
  );
};

export default DashIndex;

const Container = styled.div`
  border: solid 3px blue;
  width: 90%;
  display: flex;
  flex-direction: row;
  align-items: stretch;
`;
