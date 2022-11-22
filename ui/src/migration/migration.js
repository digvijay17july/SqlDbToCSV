import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import Alert from 'react-bootstrap/Alert';
import Table from 'react-bootstrap/Table';
function Migration(){
    return ( <Alert variant="primary" className="d-flex align-items-center justify-content-center  min-vh-100">
        
    <Container className="position-absolute top-50 start-50 translate-middle shadow p-3 mb-5 bg-body rounded">
    <Alert.Heading>Loaded Tables!</Alert.Heading>
    <p>
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>#</th>
          <th>Click To download the data</th>
        </tr>
      </thead>
            <tbody>
 
              <tr>
                <td>1</td>
                <td><a href="http://localhost:4000/api/loadData">Download</a></td>
        </tr>

      </tbody>
    </Table>
  </p>
    </Container>
    </Alert>
    )
}
export default Migration