import styled from 'styled-components';
import FeedEntry from './FeedEntry';

const Feed: React.FC = () => {
  return (
    <Container>
      Feed
      <FeedEntry></FeedEntry>
    </Container>
  );
};

export default Feed;

const Container = styled.div`
  border: 3px solid green;
  width: 30%;
`;
