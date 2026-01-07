package seed

import "github.com/pathway/backend/models"

func modulePatterns5() models.Module {
	return models.Module{
		ID:    "patterns-5",
		Title: "MVC Pattern",
		Content: []models.ContentBlock{
			textBlock(`## Model-View-Controller (MVC)

MVC is an architectural pattern that separates an application into three interconnected components.

**Components**:
- **Model**: Data and business logic
- **View**: User interface
- **Controller**: Handles input and coordinates Model and View`),
			textBlock(`## The Three Components

**Model**:
- Represents data and business logic
- Independent of UI
- Notifies observers of changes
- Handles data validation

**View**:
- Displays data to user
- Observes Model for changes
- Sends user input to Controller
- No business logic

**Controller**:
- Handles user input
- Updates Model
- Selects View to display
- Coordinates Model and View`),
			codeBlock("javascript", `// MVC Pattern Example

// Model
class UserModel {
  constructor() {
    this.users = [];
    this.observers = [];
  }
  
  subscribe(observer) {
    this.observers.push(observer);
  }
  
  notify() {
    this.observers.forEach(obs => obs.update(this.users));
  }
  
  addUser(user) {
    this.users.push(user);
    this.notify();
  }
  
  getUsers() {
    return this.users;
  }
}

// View
class UserView {
  constructor(controller) {
    this.controller = controller;
  }
  
  render(users) {
    console.log('Users:', users);
    // Render UI
  }
  
  update(users) {
    this.render(users);
  }
  
  handleAddUser(name) {
    this.controller.addUser(name);
  }
}

// Controller
class UserController {
  constructor(model) {
    this.model = model;
  }
  
  addUser(name) {
    this.model.addUser({ name, id: Date.now() });
  }
  
  setView(view) {
    this.view = view;
    this.model.subscribe(view);
  }
}

// Usage
const model = new UserModel();
const controller = new UserController(model);
const view = new UserView(controller);
controller.setView(view);

view.handleAddUser('Alice');`),
			textBlock(`## MVC Flow

1. **User interacts with View** (clicks button, fills form)
2. **View sends input to Controller**
3. **Controller updates Model**
4. **Model notifies observers** (including View)
5. **View updates display** based on new Model state`),
			calloutBlock("info", "MVC is one of the most widely used patterns. Many frameworks implement it: Ruby on Rails, Django, Angular, etc."),
			textBlock(`## Benefits of MVC

**Separation of Concerns**:
- Each component has clear responsibility
- Changes to one don't affect others

**Reusability**:
- Models can be used with different Views
- Views can display different Models

**Testability**:
- Each component can be tested independently
- Business logic separated from UI

**Maintainability**:
- Clear structure
- Easy to understand and modify`),
			textBlock(`## Variations

**MVP (Model-View-Presenter)**:
- Presenter handles all UI logic
- View is passive

**MVVM (Model-View-ViewModel)**:
- ViewModel binds View and Model
- Two-way data binding

**MVC vs MVVM**:
- MVC: Controller handles input
- MVVM: ViewModel provides data binding`),
			exerciseBlock(
				"In MVC, where should validation logic go - Model, View, or Controller?",
				"Validation logic should primarily go in the Model, because:\n1. Business rules belong in the Model\n2. Validation should be reusable across different Views/Controllers\n3. Ensures data integrity regardless of how it's entered\n\nHowever, the View can have basic input validation (format checking) for better UX, and the Controller can coordinate validation, but the Model should have the authoritative validation rules.",
				[]string{"Where does business logic belong?", "What needs to be reusable?"},
			),
		},
	}
}
