// Chapter definitions - groups modules from various courses into learning chapters
// Chapters are presented sequentially to students, with items that must be completed in order

export const chapters = [
  {
    id: 'chapter-1',
    title: 'Getting Started',
    description: 'Set up your development environment, learn Git basics, and make your first commits. By the end of this chapter, you\'ll have a working development setup and understand the basics of version control.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Git',
        moduleId: 'git-1',
        title: 'GitHub Account Setup',
        description: 'Create your GitHub account and configure Git on your machine'
      },
      {
        courseTitle: 'Git',
        moduleId: 'git-2',
        title: 'Your First Commit',
        description: 'Create a HelloWorld.txt file and make your first Git commit'
      },
      {
        courseTitle: 'Git',
        moduleId: 'git-3',
        title: 'Push to GitHub',
        description: 'Connect your local repository to GitHub and push your first code'
      },
      {
        courseTitle: 'Development Tools',
        moduleId: 'vscode-setup',
        title: 'VS Code Setup',
        description: 'Download, install, and configure Visual Studio Code'
      },
      {
        courseTitle: 'Git',
        moduleId: 'git-4',
        title: 'Branching and Merging',
        description: 'Create a development branch and learn about branching workflows'
      },
      {
        courseTitle: 'Git',
        moduleId: 'git-5',
        title: 'Your First Pull Request',
        description: 'Open a pull request and learn about code review'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'typescript-starter',
        title: 'Pull the Starter Repo',
        description: 'Clone a TypeScript repository and explore basic code concepts'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'react-hello-world',
        title: 'Your First React Page',
        description: 'Set up Vite + React and create your first web page'
      }
    ]
  },
  {
    id: 'chapter-2',
    title: 'Coding Foundations',
    description: 'Learn beginner-friendly programming fundamentals (how code is structured, how programs make decisions), then build a real TypeScript class and a fun React animation to visualize the classic FizzBuzz challenge.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Code Concepts',
        moduleId: 'code-5',
        title: 'Programming Languages 101',
        description: 'Understand how languages like C#, JavaScript, and TypeScript compare'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'code-1',
        title: 'Object-Oriented Programming',
        description: 'Learn what classes and objects are and why they matter'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'code-6',
        title: 'Control Flow',
        description: 'Learn how programs make decisions and repeat work'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'fizzbuzz-class',
        title: 'FizzBuzz Engine (TypeScript Class)',
        description: 'Build a TypeScript class that solves FizzBuzz'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'fizzbuzz-galton-ui',
        title: 'FizzBuzz Galton Board UI',
        description: 'Build a small React animation that visualizes FizzBuzz results'
      }
    ]
  },
  {
    id: 'chapter-3',
    title: 'HTTP & Web Networking',
    description: 'Learn how the web actually works: HTTP requests and responses, common methods and status codes, headers and cookies, REST API design, and the basics of HTTPS security.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-1',
        title: 'HTTP Fundamentals',
        description: 'Understand what HTTP is and how requests and responses work'
      },
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-2',
        title: 'HTTP Methods',
        description: 'Learn GET, POST, PUT, PATCH, DELETE, and when to use each'
      },
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-3',
        title: 'HTTP Status Codes',
        description: 'Learn what 200s, 400s, and 500s mean (and how to use them)'
      },
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-4',
        title: 'Headers & Cookies',
        description: 'Learn how metadata and cookies move through web requests'
      },
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-5',
        title: 'REST API Design',
        description: 'Design clean, predictable endpoints that clients can use easily'
      },
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-6',
        title: 'Request/Response Cycle',
        description: 'Understand the full lifecycle of a web request (DNS, TCP, TLS, HTTP)'
      },
      {
        courseTitle: 'HTTP Networking',
        moduleId: 'http-7',
        title: 'HTTPS & Security',
        description: 'Understand why HTTPS matters and the basics of web security'
      }
    ]
  },
  {
    id: 'chapter-4',
    title: 'Collaboration & Agile Basics',
    description: 'Level up how you work with others: handle merge conflicts confidently and learn the basics of Agile and SCRUM (roles, events, and artifacts).',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Git',
        moduleId: 'git-6',
        title: 'Resolving Conflicts',
        description: 'Learn why conflicts happen and how to resolve them safely'
      },
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-1',
        title: 'Introduction to Agile & SCRUM',
        description: 'Learn what Agile is and why teams use SCRUM'
      },
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-2',
        title: 'SCRUM Roles',
        description: 'Product Owner, Scrum Master, and the Development Team'
      },
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-3',
        title: 'SCRUM Events',
        description: 'Sprints, standups, reviews, and retrospectives'
      },
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-4',
        title: 'SCRUM Artifacts',
        description: 'Product backlog, sprint backlog, and increment'
      }
    ]
  },
  {
    id: 'chapter-5',
    title: 'Shipping as a Team',
    description: 'Turn planning into delivery: learn sprint planning, estimation, and how teams actually ship work. Then reinforce it with advanced Git workflows used in real teams.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-5',
        title: 'Sprint Planning',
        description: 'Plan a sprint and commit to a realistic goal'
      },
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-6',
        title: 'User Stories & Estimation',
        description: 'Write clear stories and estimate work as a team'
      },
      {
        courseTitle: 'SCRUM',
        moduleId: 'scrum-7',
        title: 'Running Effective Sprints',
        description: 'Practical tips for staying on track and delivering value'
      },
      {
        courseTitle: 'Git',
        moduleId: 'git-7',
        title: 'Advanced Git',
        description: 'Learn rebase, cherry-pick, and stash (when you actually need them)'
      }
    ]
  },
  {
    id: 'chapter-6',
    title: 'Code Building Blocks',
    description: 'Go beyond “it works” and learn patterns of thinking that make code easier to read, reuse, and reason about.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Code Concepts',
        moduleId: 'code-4',
        title: 'Functions & Closures',
        description: 'Learn how functions can take inputs, return outputs, and capture values'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'code-2',
        title: 'Functional Programming',
        description: 'Learn a different style of writing code with functions and data flows'
      },
      {
        courseTitle: 'Code Concepts',
        moduleId: 'code-3',
        title: 'Protocol Oriented Programming',
        description: 'Learn how “contracts” (interfaces/protocols) help structure code'
      }
    ]
  },
  {
    id: 'chapter-7',
    title: 'SOLID Foundations (Maintainable Code)',
    description: 'Learn the SOLID principles so your code stays clean as it grows. This is where you start thinking like an engineer, not just a coder.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-1',
        title: 'Introduction to SOLID Principles',
        description: 'What SOLID is and why it matters'
      },
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-2',
        title: 'Single Responsibility Principle (SRP)',
        description: 'One class, one job'
      },
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-3',
        title: 'Open/Closed Principle (OCP)',
        description: 'Add new behavior without rewriting old code'
      },
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-4',
        title: 'Liskov Substitution Principle (LSP)',
        description: 'Subtypes should behave like their base types'
      },
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-5',
        title: 'Interface Segregation Principle (ISP)',
        description: 'Small, focused interfaces are easier to use'
      },
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-6',
        title: 'Dependency Inversion Principle (DIP)',
        description: 'Depend on abstractions, not concrete details'
      },
      {
        courseTitle: 'Architecture - SOLID',
        moduleId: 'solid-7',
        title: 'Applying SOLID in Practice',
        description: 'How SOLID shows up in real codebases'
      }
    ]
  },
  {
    id: 'chapter-8',
    title: 'Design Patterns (Practical Reuse)',
    description: 'Learn proven design patterns that help you solve common problems in clean, reusable ways.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-1',
        title: 'Introduction to Design Patterns',
        description: 'Why patterns exist and how to think about them'
      },
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-2',
        title: 'Creational Patterns',
        description: 'Singleton, Factory, and object creation strategies'
      },
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-3',
        title: 'Structural Patterns',
        description: 'Adapter, Decorator, and combining objects safely'
      },
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-4',
        title: 'Behavioral Patterns',
        description: 'Observer, Strategy, and how objects communicate'
      },
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-5',
        title: 'MVC Pattern',
        description: 'A common way to structure apps (Model-View-Controller)'
      },
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-6',
        title: 'Repository Pattern',
        description: 'Separate business logic from data access'
      },
      {
        courseTitle: 'Design Patterns',
        moduleId: 'patterns-7',
        title: 'Choosing the Right Pattern',
        description: 'How to pick patterns without overengineering'
      }
    ]
  },
  {
    id: 'chapter-9',
    title: 'Testing & Quality (Confidence to Ship)',
    description: 'Learn how to prove your code works, catch bugs early, and ship changes without fear.',
    estimatedTime: '1 week',
    items: [
      {
        courseTitle: 'Testing',
        moduleId: 'testing-1',
        title: 'Testing Fundamentals',
        description: 'Why tests matter and what types of tests exist'
      },
      {
        courseTitle: 'Testing',
        moduleId: 'testing-2',
        title: 'Unit Testing',
        description: 'Test small pieces of code quickly'
      },
      {
        courseTitle: 'Testing',
        moduleId: 'testing-3',
        title: 'Integration Testing',
        description: 'Test how parts work together'
      },
      {
        courseTitle: 'Testing',
        moduleId: 'testing-4',
        title: 'Test-Driven Development (TDD)',
        description: 'Write tests first, then write the code'
      },
      {
        courseTitle: 'Testing',
        moduleId: 'testing-5',
        title: 'Mocking & Stubbing',
        description: 'Replace dependencies so you can test reliably'
      },
      {
        courseTitle: 'Testing',
        moduleId: 'testing-6',
        title: 'Testing Best Practices',
        description: 'Write tests that stay useful over time'
      },
      {
        courseTitle: 'Testing',
        moduleId: 'testing-7',
        title: 'Continuous Testing',
        description: 'Keep quality high with automation (CI)'
      }
    ]
  }
];

// Helper function to get a chapter by ID
export const getChapterById = (chapterId) => {
  return chapters.find(chapter => chapter.id === chapterId);
};

// Helper function to get all chapter IDs
export const getChapterIds = () => {
  return chapters.map(chapter => chapter.id);
};

