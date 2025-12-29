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

