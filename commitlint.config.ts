// commitLintConfig rules description:
// 'init'       - Initial commit or setup for the project.
// 'build'      - Changes that affect the build system or external dependencies (example: npm, webpack, etc.).
// 'chore'      - Regular maintenance or minor task that doesn't affect production code (example: updating dependencies).
// 'ci'         - Changes to CI configuration files and scripts (example: GitHub Actions, CircleCI, etc.).
// 'docs'       - Documentation changes or improvements.
// 'setup'      - Configuration or setup related changes (example: project setup, tooling configuration).
// 'feat'       - A new feature or functionality.
// 'fix'        - A bug fix.
// 'hot-fix'    - A critical or urgent bug fix.
// 'seo'        - SEO-related improvements or updates.
// 'story'      - Changes related to stories (example: updating storybook, UI stories).
// 'upgrade'    - Dependency upgrade or version bump.
// 'downgrade'  - Dependency downgrade or version rollback.
// 'perf'       - A code change that improves performance.
// 'refactor'   - Code changes that neither fix a bug nor add a feature, but improve the code structure.
// 'revert'     - Reverts a previous commit.
// 'style'      - Code style updates (example: formatting, missing semi-colons, etc. that do not change functionality).
// 'test'       - Adding or updating tests.
// 'design'     - Design changes or updates (example: UI/UX improvements).
// 'translation'- Changes related to localization or translations.
// 'security'   - Security improvements or patches.
// 'changeset'  - Change log or version bump-related changes.

const commitLintConfig = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'body-leading-blank': [1, 'always'],
    'body-max-line-length': [2, 'always', 100],
    'footer-leading-blank': [1, 'always'],
    'footer-max-line-length': [2, 'always', 100],
    'header-max-length': [2, 'always', 100],
    'scope-case': [2, 'always', 'lower-case'],
    'subject-case': [
      2,
      'never',
      ['sentence-case', 'start-case', 'pascal-case', 'upper-case'],
    ],
    'subject-empty': [2, 'never'],
    'subject-full-stop': [2, 'never', '.'],
    'type-case': [2, 'always', 'lower-case'],
    'type-empty': [2, 'never'],
    'type-enum': [
      2,
      'always',
      [
        'init',
        'build',
        'chore',
        'ci',
        'docs',
        'setup',
        'feat',
        'fix',
        'hot-fix',
        'seo',
        'story',
        'upgrade',
        'downgrade',
        'perf',
        'refactor',
        'revert',
        'style',
        'test',
        'design',
        'translation',
        'security',
        'changeset',
      ],
    ],
  },
};

export default commitLintConfig;
