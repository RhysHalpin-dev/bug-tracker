import React from 'react';
import styled from 'styled-components';

interface FeedDetails {
  user?: string;
  action?: string;
  project?: string;
}

const FeedEntry: React.FC<FeedDetails> = () => {
  return (
    <Container>
      <header>
        <div className="col"> User</div>
        <div className="col"> Action</div>
        <div className="col"> Project</div>
      </header>
      <div className="row">
        <div className="col"> 1</div>
        <div className="col"> 2</div>
        <div className="col"> 3</div>
      </div>
    </Container>
  );
};

export default FeedEntry;

const Container = styled.div`
  border: solid 3px black;

  height: auto;

  header,
  .row {
    display: flex;
  }
  .col {
    flex: 1;
  }
`;
