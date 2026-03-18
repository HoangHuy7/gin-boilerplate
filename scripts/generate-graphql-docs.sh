#!/bin/bash

# Script generate GraphQL documentation using SpectaQL
# Usage: ./generate-graphql-docs.sh

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
GRAPHQL_ENDPOINT="http://localhost:8082/graphql"
OUTPUT_DIR="./docs/graphql"
CONFIG_FILE="./scripts/spectaql-config.yml"

echo -e "${GREEN}🚀 Generating GraphQL Documentation...${NC}"

# Check if spectaql is installed
if ! command -v spectaql &> /dev/null; then
    echo -e "${YELLOW}⚠️  SpectaQL not found. Installing...${NC}"
    
    # Check if npm is installed
    if ! command -v npm &> /dev/null; then
        echo -e "${RED}❌ npm is not installed. Please install Node.js first.${NC}"
        echo "   Visit: https://nodejs.org/"
        exit 1
    fi
    
    # Install spectaql globally
    npm install -g spectaql
    echo -e "${GREEN}✅ SpectaQL installed successfully${NC}"
fi

# Create output directory if not exists
mkdir -p "$OUTPUT_DIR"

# Check if GraphQL server is running
echo -e "${YELLOW}🔍 Checking GraphQL endpoint: $GRAPHQL_ENDPOINT${NC}"
if ! curl -s "$GRAPHQL_ENDPOINT" > /dev/null; then
    echo -e "${RED}❌ GraphQL server is not running at $GRAPHQL_ENDPOINT${NC}"
    echo "   Please start your Go server first: go run ./cmd/gas/"
    exit 1
fi

echo -e "${GREEN}✅ GraphQL server is running${NC}"

# Generate documentation
echo -e "${YELLOW}📄 Generating documentation...${NC}"
spectaql "$CONFIG_FILE" -t "$OUTPUT_DIR"

echo -e "${GREEN}✅ Documentation generated successfully!${NC}"
echo -e "${GREEN}📁 Output: $OUTPUT_DIR/index.html${NC}"
echo -e "${YELLOW}🌐 Open file://$(pwd)/$OUTPUT_DIR/index.html in your browser${NC}"
