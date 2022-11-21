import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import Alert from 'react-bootstrap/Alert';
function Welcome(){

    return (  
      <Alert variant="primary" className="d-flex align-items-center justify-content-center  min-vh-100">
        
      <Container className="position-absolute top-50 start-50 translate-middle shadow p-3 mb-5 bg-body rounded">
      <Alert.Heading>Settings for DB Connection!</Alert.Heading>
      <p>
      <Form action='/migration'>
      <Form.Group className="mb-3" controlId="formBasicUsernameID">
        <Form.Label >Username</Form.Label>
        <Form.Control type="text" placeholder="Enter username" />
        <Form.Text className="text-muted">
          Please Enter the DB Username
        </Form.Text>
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" />
        <Form.Text className="text-muted">
          Please Enter the DB Password
        </Form.Text>
      </Form.Group>
      <Form.Group className="mb-3" controlId="formBasicDbIP">
        <Form.Label>DB IP ADDRESS</Form.Label>
        <Form.Control type="text" placeholder="Enter DB IP Address" />
        <Form.Text className="text-muted">
          Please Enter the DB IP ADDRESS
        </Form.Text>
      </Form.Group>
      <Button variant="primary" type="submit">
        NEXT
      </Button>
    </Form>
    </p>
      </Container>
      </Alert>
      )
}

export default Welcome