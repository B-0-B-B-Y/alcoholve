import { render, screen } from '@testing-library/react';
import { App } from '../components/App';

test('renders welcome to alcoholve title', () => {
  render(<App />);
  const titleElement = screen.getByText(/alcoholve/i);
  expect(titleElement).toBeInTheDocument();
});
