package language_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestYAMLSimpleKeyValueLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Simple Key Value Language Detection")
	code := `key: value
another_key: Another value goes here.
a_number_value: 100
scientific_notation: 1e+12`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLCollectionLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Collection Language Detection")
	code := `# Nesting uses indentation. 2 space indent is preferred (but not required).
a_nested_map:
  key: value
  another_key: Another Value
  another_nested_map:
    hello: hello

# Maps don't have to have string keys.
0.25: a float key`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLComplexKeyLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Complex Key Language Detection")
	code := `? |
  This is a key
  that has multiple lines
: and this is its value`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYamlMergeKeyLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Merge Key Language Detection")
	code := `foo:
  <<: *base
  age: 10

bar:
  <<: *base
  age: 20`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLBinaryLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Binary Language Detection")
	code := `gif_file: !!binary |
   R0lGODlhDAAMAIQAAP//9/X17unp5WZmZgAAAOfn515eXvPz7Y6OjuDg4J+fn5
   OTk6enp56enmlpaWNjY6Ojo4SEhP/++f/++f/++f/++f/++f/++f/++f/++f/+
   +f/++f/++f/++f/++f/++SH+Dk1hZGUgd2l0aCBHSU1QACwAAAAADAAMAAAFLC
   AgjoEwnuNAFOhpEMTRiggcz4BNJHrv/zCFcLiwMWYNG84BwwEeECcgggoBADs=`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLSetTypesLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Set Types Language Detection")
	code := `set:
  ? item1
  ? item2
  ? item3
or: {item1, item2, item3}`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLDockercomposeFileLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Dockercompose File Language Detection")
	code := `version: "3.8"
  services:
    dynamodb-local:
      image: amazon/dynamodb-local:1.16.0
      container_name: dynamodb-local
      ports:
        - "8000:8000"
      volumes:
        - "./docker/dynamodb:/home/dynamodblocal/data"
      working_dir: /home/dynamodblocal
      command: "-jar DynamoDBLocal.jar -sharedDb -optimizeDbBeforeStartup -dbPath ./data"`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLGithubActionsLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Github Actions Language Detection")
	code := `n:
  release:
    types: [created]

env:
  AZURE_WEBAPP_NAME: your-app-name    # set this to your application's name
  AZURE_WEBAPP_PACKAGE_PATH: '.'      # set this to the path to your web app project, defaults to the repository root
  NODE_VERSION: '10.x'                # set this to the node version to use

jobs:
  build-and-deploy:
    name: Build and Deploy
    runs-on: ubuntu-latest
    environment: production
    steps:
    - uses: actions/checkout@v2
    - name: Use Node.js \${{ env.NODE_VERSION }}
      uses: actions/setup-node@v2
      with:
        node-version: \${{ env.NODE_VERSION }}
    - name: npm install, build, and test
      run: |
        # Build and test the project, then
        # deploy to Azure Web App.
        npm install
        npm run build --if-present
        npm run test --if-present
    - name: 'Deploy to Azure WebApp'
      uses: azure/webapps-deploy@v2
      with:
        app-name: \${{ env.AZURE_WEBAPP_NAME }}
        publish-profile: \${{ secrets.AZURE_WEBAPP_PUBLISH_PROFILE }}
        package: \${{ env.AZURE_WEBAPP_PACKAGE_PATH }}`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLCircleCIConfigLanguageDetection(t *testing.T) {
	t.Log("Testing YAML CircleCI Config Language Detection")
	code := `aliases:
  # Cache management
  - &restore_yarn_cache
    restore_cache:
      keys:
        - v1-yarn-cache

  - &save_yarn_cache
    save_cache:
      paths:
        - ~/.cache/yarn
      key: v1-yarn-cache

  - &restore_deps_cache
    restore_cache:
      keys:
        - v1-deps-cache-{{ checksum "yarn.lock" }}

  - &save_deps_cache
    save_cache:
      paths:
        - node_modules
      key: v1-yarn-deps-{{ checksum "yarn.lock" }}

  # Default
  - &defaults
    working_directory: ~/prettier
    docker:
      - image: circleci/node:9

version: 2
jobs:
  # Install dependencies and cache everything
  checkout_code:
    <<: *defaults
    steps:
      - checkout
      - *restore_yarn_cache
      - *restore_deps_cache
      - run: yarn install
      - run: yarn check-deps
      - *save_deps_cache
      - *save_yarn_cache
      - persist_to_workspace:
          root: .
          paths:
            - .

  # Create the production bundle and cache
  build_prod:
    <<: *defaults
    environment:
      NODE_ENV: production
    steps:
      - attach_workspace:
          at: ~/prettier
      - run: yarn build
      - persist_to_workspace:
          root: .
          paths:
            - dist
      - store_artifacts:
          path: ~/prettier/dist

  # Run tests on the production bundle
  test_prod_node4:
    <<: *defaults
    docker:
      - image: circleci/node:4
    steps:
      - attach_workspace:
          at: ~/prettier
      - run: yarn test:dist

  # Run tests on the production bundle
  test_prod_node9:
    <<: *defaults
    steps:
      - attach_workspace:
          at: ~/prettier
      - run: yarn test:dist

workflows:
  version: 2
  prod:
    jobs:
      - checkout_code
      - build_prod:
          requires:
            - checkout_code
      - test_prod_node4:
          requires:
            - build_prod
      - test_prod_node9:
          requires:
            - build_prod`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestYAMLEslintConfigLanguageDetection(t *testing.T) {
	t.Log("Testing YAML Eslint Config Language Detection")
	code := `
  root: true
  parser: '@typescript-eslint/parser'
  plugins:
    - '@typescript-eslint'
  extends:
    - eslint:recommended
    #  - plugin:react/recommended
    - plugin:@typescript-eslint/recommended
    - prettier
    - prettier/@typescript-eslint
  env:
    node: true
    browser: true
    es6: true
    jest: true
  rules:
    '@typescript-eslint/explicit-function-return-type': off
    '@typescript-eslint/explicit-member-accessibility': off
    '@typescript-eslint/indent': off
    '@typescript-eslint/member-delimiter-style': off
    '@typescript-eslint/no-explicit-any': off
    '@typescript-eslint/no-empty-function': off
    '@typescript-eslint/no-non-null-assertion': off
    '@typescript-eslint/no-var-requires': off
    '@typescript-eslint/explicit-module-boundary-types': off
    '@typescript-eslint/ban-ts-comment': off
    '@typescript-eslint/ban-types':
      - 2
      - types:
          Function: false
    '@typescript-eslint/no-unused-vars':
      - 2
      - argsIgnorePattern: '^_'`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_YAML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_YAML, result.Language)
		t.Errorf("Result: %v", result)
	}
}
