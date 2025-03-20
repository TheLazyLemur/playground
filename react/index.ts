import { createRoot } from 'react-dom/client';
import { Body } from './components';

const contentRoot = document.getElementById('react-content');
if (!contentRoot) {
    throw new Error('Could not find element with id react-content');
}
const contentReactRoot = createRoot(contentRoot);
contentReactRoot.render(Body());

